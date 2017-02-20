package main

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
	"github.com/upwork/golang-upwork/api/routers/metadata"
	"gopkg.in/cheggaaa/pb.v1"
)

func skills(settings *Settings) {
	skills := get_skills(settings)
	skills_count := len(skills)

	database := get_database(settings)

	progress_bar := pb.StartNew(skills_count)
	for _, skill := range skills {
		skills_insert(database, skill)
		progress_bar.Increment()
	}
	progress_bar.Finish()

	titles := []string{}
	for _, skill := range skills {
		titles = append(titles, skill.Title)
	}
	skills_prune(database, titles)
}

func get_skills(settings *Settings) []Skill {
	client := get_client(settings)
	_, bytes := metadata.New(client).GetSkills()
	upwork_skills := UpworkSkills{}
	err := json.Unmarshal(bytes, &upwork_skills)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	skills := []Skill{}
	for _, title := range upwork_skills.Items {
		skill := Skill{
			Title:  title,
			Status: "On",
		}
		skills = append(skills, skill)
	}
	return skills
}
