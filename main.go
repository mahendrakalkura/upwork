package main

import (
	"flag"
	"github.com/getsentry/raven-go"
)

func main() {
	settings := get_settings()

	raven.SetDSN(settings.Raven.Dsn)

	action := flag.String("action", "", "")
	flag.Parse()
	if *action == "bootstrap" {
		bootstrap(settings)
	}
	if *action == "categories" {
		categories(settings)
	}
	if *action == "jobs" {
		jobs(settings)
	}
	if *action == "skills" {
		skills(settings)
	}
	if *action == "ui" {
		ui(settings)
	}
}
