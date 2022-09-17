package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4107728804177-4095022982163-GwHv1mhrLKXIZAEL95s2hLZV")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A042Q49GJJH-4088498220070-990f57ae0337ca36016bc394deafd0823ce2dc148eb4a8f16828c60eb347f5a2")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	var strSample []string
	strSample = append(strSample, "!nascimento")
	bot.Command("nasci em <nascimento>", &slacker.CommandDefinition{
		Description: "calculadora de idade",
		Examples:    strSample,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("nascimento")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("você tem %d anos", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
