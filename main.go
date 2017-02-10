package main

import (
	"bufio"
	"fmt"
	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/auth"
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
	_, json := auth.New(client).GetUserInfo()
	fmt.Println(string(json))
}
