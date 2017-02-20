package main

import (
	"fmt"
	"github.com/upwork/golang-upwork/api/routers/jobs/search"
)

func jobs(settings *Settings) {
	client := get_client(settings)
	params := map[string]string{
		"q": "",
	}
	_, bytes := search.New(client).Find(params)
	fmt.Println(string(bytes))
}
