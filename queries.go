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
