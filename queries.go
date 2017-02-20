package main

import (
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
)

func categories_insert(database *sqlx.DB, category Category) {
	statement := `SELECT categories_insert (:title, :sub_title, :status)`
	database.NamedExec(statement, category)
}

func categories_prune(database *sqlx.DB, titles_and_sub_titles []string) {
	query := `DELETE FROM categories WHERE CONCAT(title, sub_title) NOT IN (?)`
	statement, arguments, err := sqlx.In(query, titles_and_sub_titles)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	statement = database.Rebind(statement)
	database.Exec(query, arguments)
}

func jobs_insert(database *sqlx.DB, job Job) {
	statement := `
	SELECT jobs_insert (
        :budget,
        :category,
        :client_country,
        :client_feedback,
        :client_jobs_posted,
        :client_past_hires,
        :client_reviews_count,
        :date_created,
        :duration,
        :job_status,
        :job_type,
        :skills,
        :snippet,
        :sub_category,
        :title,
        :url,
        :workload
	)
    `
	database.NamedExec(statement, job)
}

func jobs_prune(database *sqlx.DB, urls []string) {
	query := `DELETE FROM jobs WHERE url NOT IN (?)`
	statement, arguments, err := sqlx.In(query, urls)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	statement = database.Rebind(statement)
	database.Exec(query, arguments)
}

func skills_insert(database *sqlx.DB, skill Skill) {
	statement := `SELECT skills_insert (:title, :status)`
	database.NamedExec(statement, skill)
}

func skills_prune(database *sqlx.DB, titles []string) {
	query := `DELETE FROM skills WHERE title NOT IN (?)`
	statement, arguments, err := sqlx.In(query, titles)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	statement = database.Rebind(statement)
	database.Exec(query, arguments)
}
