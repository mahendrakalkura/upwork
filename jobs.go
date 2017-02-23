package main

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	"github.com/upwork/golang-upwork/api/routers/jobs/search"
	"gopkg.in/cheggaaa/pb.v1"
	"strconv"
	"sync"
)

var wait_group sync.WaitGroup

func jobs(settings *Settings) {
	database := get_database(settings)

	channels_upwork_jobs_requests := make(chan UpworkJobsRequest, settings.Others.Consumers*2)
	channels_jobs := make(chan []Job, settings.Others.Consumers*2)

	wait_group.Add(settings.Others.Consumers)
	for index := 1; index <= settings.Others.Consumers; index++ {
		go jobs_consumer(settings, database, channels_upwork_jobs_requests, channels_jobs)
	}

	wait_group.Add(1)
	go jobs_producer(database, channels_upwork_jobs_requests)

	go func() {
		wait_group.Wait()
		close(channels_jobs)
	}()

	urls := []string{}

	for jobs := range channels_jobs {
		jobs_count := len(jobs)
		progress_bar := pb.StartNew(jobs_count)
		for _, job := range jobs {
			urls = append(urls, job.Url)
			jobs_insert(database, job)
			progress_bar.Increment()
		}
		progress_bar.Finish()
	}

	if len(urls) > 0 {
		jobs_prune(database, urls)
	}
}

func jobs_consumer(
	settings *Settings,
	database *sqlx.DB,
	channels_upwork_jobs_requests chan UpworkJobsRequest,
	channels_jobs chan []Job,
) {
	defer wait_group.Done()

	for upwork_jobs_request := range channels_upwork_jobs_requests {
		jobs := get_jobs(settings, upwork_jobs_request)
		channels_jobs <- jobs
	}
}

func jobs_producer(database *sqlx.DB, channels_upwork_jobs_requests chan UpworkJobsRequest) {
	defer wait_group.Done()

	upwork_jobs_requests := get_upwork_jobs_requests()
	for _, upwork_jobs_request := range upwork_jobs_requests {
		channels_upwork_jobs_requests <- upwork_jobs_request
	}

	close(channels_upwork_jobs_requests)
}

func get_upwork_jobs_requests() []UpworkJobsRequest {
	job_types := []string{"hourly", "fixed-price"}
	durations := []string{"month", "quarter", "semester", "ongoing"}
	workloads := []string{"full_time"}
	upwork_jobs_requests := []UpworkJobsRequest{}
	for _, job_type := range job_types {
		for _, duration := range durations {
			for _, workload := range workloads {
				upwork_jobs_requests = append(
					upwork_jobs_requests,
					UpworkJobsRequest{
						Count:      100,
						DaysPosted: 1,
						Duration:   duration,
						JobStatus:  "open",
						JobType:    job_type,
						Offset:     0,
						Q:          "",
						Workload:   workload,
					},
				)
			}
		}
	}
	return upwork_jobs_requests
}

func get_jobs(settings *Settings, upwork_jobs_request UpworkJobsRequest) []Job {
	jobs := []Job{}
	client := get_client(settings)
	for {
		paging := get_paging(upwork_jobs_request.Offset, upwork_jobs_request.Count)
		parameters := map[string]string{
			"days_posted": strconv.FormatInt(upwork_jobs_request.DaysPosted, 10),
			"duration":    upwork_jobs_request.Duration,
			"job_status":  upwork_jobs_request.JobStatus,
			"job_type":    upwork_jobs_request.JobType,
			"paging":      paging,
			"q":           upwork_jobs_request.Q,
			"workload":    upwork_jobs_request.Workload,
		}
		_, bytes := search.New(client).Find(parameters)
		upwork_jobs_response := UpworkJobsResponse{}
		err := json.Unmarshal(bytes, &upwork_jobs_response)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		}
		jobs = append(jobs, upwork_jobs_response.Jobs...)
		offset_count := upwork_jobs_response.Paging.Offset + upwork_jobs_response.Paging.Count
		if offset_count >= upwork_jobs_response.Paging.Total {
			break
		}
		upwork_jobs_request.Offset = offset_count
	}
	return jobs
}
