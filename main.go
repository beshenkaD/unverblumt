package main

import (
	"flag"
	"log"
	"os"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/cat"
	"github.com/beshenkaD/unverblumt/hru"
	"github.com/beshenkaD/unverblumt/me"
	"github.com/beshenkaD/unverblumt/openweathermap"
	"github.com/beshenkaD/unverblumt/ping"
	"github.com/beshenkaD/unverblumt/quote"
	"github.com/beshenkaD/unverblumt/shortener"
)

func main() {
	version := "0.3.2"

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
			{Name: "dark | light", Desc: "background color", Optional: true},
		},
		quote.Quote)

	b.RegisterCommand(
		"/weather",
		"Displays current weather",
		[]bot.CommandParam{
			{Name: "<city>", Desc: "get weather for <city>. Default value is: \"Moscow\"", Optional: true},
			{Name: "<lang>", Desc: "language code. Default value is: \"en\"", Optional: true},
		},
		openweathermap.Weather)

	b.RegisterCommand(
		"/forecast",
		"Displays forecast",
		[]bot.CommandParam{
			{Name: "<city>", Desc: "get weather for <city>. Default value is: \"Moscow\"", Optional: true},
			{Name: "<lang>", Desc: "language code. Default value is: \"en\"", Optional: true},
		},
		openweathermap.Forecast)

	b.RegisterCommand(
		"/hru",
		"Pig fun funny pig ahaahhah hruhruhru",
		nil,
		hru.HruCommand)

	b.RegisterCommand(
		"/cat",
		"Sends random cat image. How cute!",
		nil,
		cat.Cat)

	b.RegisterCommand(
		"/short",
		"Shorts given URL",
		[]bot.CommandParam{
			{Name: "<target>", Desc: "URL which will be shortened", Optional: false},
			{Name: "<Expire in>", Desc: "After this time, the link will stop working. Format: 2[m|h|d] (minutes, hours, days)", Optional: true},
			{Name: "<Password>", Desc: "Guess what it is", Optional: true},
		},
		shortener.Short)

	b.RegisterCommand(
		"/shortBunch",
		"Shorts multiple URLs",
		[]bot.CommandParam{
			{Name: "string", Desc: "this command use \"<target>:<ixpireIn>:<password>\" format.\nIf you want to ignore param just do it. Example (ignore expireIn):\n/shortBunch \"example.com::password\"", Optional: false},
		},
		shortener.ShortBunch)

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
