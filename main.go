package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/jobs/search"
	"github.com/upwork/golang-upwork/api/routers/metadata"
	"os"
)

func main() {
	config := api.ReadConfig("config.json")
	client := api.Setup(config)
	if !client.HasAccessToken() {
		url := client.GetAuthorizationUrl("")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(url)
		verifier, _ := reader.ReadString('\n')
		token := client.GetAccessToken(verifier)
		fmt.Println(token)
	}

	action := flag.String("action", "", "")
	flag.Parse()
	if *action == "categories" {
		categories(client)
	}
	if *action == "jobs" {
		jobs(client)
	}

}

func categories(client api.ApiClient) {
	_, json := metadata.New(client).GetCategoriesV2()
	fmt.Println(string(json))
}

func jobs(client api.ApiClient) {
	params := map[string]string{
		"q": "",
	}
	_, json := search.New(client).Find(params)
	fmt.Println(string(json))
}
