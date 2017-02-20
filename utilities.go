package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/getsentry/raven-go"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/upwork/golang-upwork/api"
	"os"
)

func get_client(settings *Settings) api.ApiClient {
	settings_upwork := map[string]string{
		"access_token":    settings.Upwork.AccessToken,
		"access_secret":   settings.Upwork.AccessSecret,
		"consumer_key":    settings.Upwork.ConsumerKey,
		"consumer_secret": settings.Upwork.ConsumerSecret,
	}
	config := api.NewConfig(settings_upwork)
	client := api.Setup(config)
	if !client.HasAccessToken() {
		url := client.GetAuthorizationUrl("")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(url)
		verifier, _ := reader.ReadString('\n')
		token := client.GetAccessToken(verifier)
		fmt.Println(token)
	}
	return client
}

func get_database(settings *Settings) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.SQLX.Hostname,
		settings.SQLX.Port,
		settings.SQLX.Username,
		settings.SQLX.Password,
		settings.SQLX.Database,
	)
	database := sqlx.MustConnect("postgres", dsn)
	return database
}

func get_settings() *Settings {
	var settings = &Settings{}
	_, err := toml.DecodeFile("settings.toml", settings)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	return settings
}
