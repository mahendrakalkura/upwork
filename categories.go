package main

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/upwork/golang-upwork/api/routers/metadata"
	"gopkg.in/cheggaaa/pb.v1"
)

func categories(settings *Settings) {
	categories := get_categories(settings)
	categories_count := len(categories)

	database := get_database(settings)

	progress_bar := pb.StartNew(categories_count)
	for _, category := range categories {
		categories_insert(database, category)
		progress_bar.Increment()
	}
	progress_bar.Finish()

	titles_and_sub_titles := []string{}
	for _, category := range categories {
		title_and_sub_title := fmt.Sprintf(
			"%s%s", category.Title, category.SubTitle,
		)
		titles_and_sub_titles = append(
			titles_and_sub_titles, title_and_sub_title,
		)
	}

	categories_prune(database, titles_and_sub_titles)
}

func get_categories(settings *Settings) []Category {
	client := get_client(settings)
	_, bytes := metadata.New(client).GetCategoriesV2()
	upwork_categories := UpworkCategories{}
	err := json.Unmarshal(bytes, &upwork_categories)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	categories := []Category{}
	for _, parent := range upwork_categories.Items {
		for _, child := range parent.Topics {
			category := Category{
				Title:    parent.Title,
				SubTitle: child.Title,
				Status:   "On",
			}
			categories = append(categories, category)
		}
	}
	return categories
}
