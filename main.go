package main

import (
	"flag"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/me"
	"github.com/beshenkaD/unverblumt/openweathermap"
	"github.com/beshenkaD/unverblumt/quote"
)

func main() {
	version := "0.1.3"

	tokFlag := flag.String("token", "", "telegram bot token")
	debFlag := flag.Bool("debug", false, "use debug")

	flag.Parse()

	b := bot.New(*tokFlag, version, *debFlag)
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
		nil,
		openweathermap.Weather)

	b.RegisterCommand(
		"/forecast",
		"Get forecast for 5 days",
		nil,
		openweathermap.Forecast)
	b.Run()
}
