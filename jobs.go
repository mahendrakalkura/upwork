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
	params := map[string]string{
		"job_status": "open",
		"paging":     "0;100",
		"q":          "",
	}
	_, bytes := search.New(client).Find(params)
	upwork_jobs := UpworkJobs{}
	err := json.Unmarshal(bytes, &upwork_jobs)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	return upwork_jobs.Jobs
}
