package main

import (
	"flag"
	"log"
	"os"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/hru"
	"github.com/beshenkaD/unverblumt/me"
	"github.com/beshenkaD/unverblumt/openweathermap"
	"github.com/beshenkaD/unverblumt/ping"
	"github.com/beshenkaD/unverblumt/quote"
)

func main() {
	version := "0.2.6"

	debFlag := flag.Bool("debug", false, "use debug")
	comFlag := flag.Bool("commands", false, "print commands for BotFather")

	flag.Parse()

	token := os.Getenv("BOT_TG")
	if token == "" {
		log.Fatalln("please export token as BOT_TG")
	}

	b := bot.New(token, version, *debFlag)
	b.RegisterCommand(
		"/ping",
		"Checks bot availability",
		nil,
		ping.Ping)

	b.RegisterCommand(
		"/me",
		"Displays a message about yourself",
		nil,
		me.Me)

	b.RegisterCommand(
		"/quote",
		"Quotes a message",
		[]bot.CommandParam{
			{Name: "dark", Desc: "use black background", Optional: true},
			{Name: "light", Desc: "use light background", Optional: true},
		},
		quote.Quote)

	b.RegisterCommand(
		"/weather",
		"Get current weather",
		[]bot.CommandParam{
			{Name: "<city>", Desc: "get weather for <city>. Default value is: \"Moscow\"", Optional: true},
			{Name: "<lang>", Desc: "language code. Default value is: \"en\"", Optional: true},
		},
		openweathermap.Weather)

	b.RegisterCommand(
		"/hru",
		"Pig fun funny pig ahaahhah hruhruhru",
		nil,
		hru.HruCommand)

	b.RegisterHook(
		"hru",
		"Detects hru and do fun stuff ahhaha hruhru",
		hru.HruWatcher)

	if *comFlag {
		b.PrintCommands()
		return
	}

	b.Run()
}
