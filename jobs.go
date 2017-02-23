package main

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
	"github.com/upwork/golang-upwork/api/routers/jobs/search"
	"gopkg.in/cheggaaa/pb.v1"
)

func jobs(settings *Settings) {
	jobs := get_jobs(settings)
	jobs_count := len(jobs)

	database := get_database(settings)

	progress_bar := pb.StartNew(jobs_count)
	for _, job := range jobs {
		jobs_insert(database, job)
		progress_bar.Increment()
	}
	progress_bar.Finish()

	urls := []string{}
	for _, job := range jobs {
		urls = append(urls, job.Url)
	}
	if len(urls) > 0 {
		jobs_prune(database, urls)
	}
}

func get_jobs(settings *Settings) []Job {
	client := get_client(settings)
	parameters := get_parameters()
	jobs := []Job{}
	parameters_count := len(parameters)
	progress_bar := pb.StartNew(parameters_count)
	for _, params := range parameters {
		params := map[string]string{
			"days_posted": "1",
			"duration":    params.Duration,
			"job_status":  "open",
			"job_type":    params.JobType,
			"paging":      "0;100",
			"q":           "",
			"workload":    params.Workload,
		}
		_, bytes := search.New(client).Find(params)
		upwork_jobs := UpworkJobs{}
		err := json.Unmarshal(bytes, &upwork_jobs)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		}
		jobs = append(jobs, upwork_jobs.Jobs...)
		progress_bar.Increment()
	}
	progress_bar.Finish()
	return jobs
}

func get_parameters() []Parameters {
	job_types := []string{"hourly", "fixed-price"}
	durations := []string{"month", "quarter", "semester", "ongoing"}
	workloads := []string{"full_time"}
	parameters := []Parameters{}
	for _, job_type := range job_types {
		for _, duration := range durations {
			for _, workload := range workloads {
				parameters = append(
					parameters,
					Parameters{
						JobType:  job_type,
						Duration: duration,
						Workload: workload,
					},
				)
			}
		}
	}
	return parameters
}
