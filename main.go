package main

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {
	slackBotToken := viperEnvVariable("SLACK_BOT_TOKEN")
	slackChannelID := viperEnvVariable("CHANNEL_ID")

	api := slack.New(slackBotToken)
	channelArr := []string{slackChannelID}

	fileArr := []string{"test.pdf"}

	for _, v := range fileArr {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     v,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URLPrivate)
	}
}
