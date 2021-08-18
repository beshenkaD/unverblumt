package main

import (
	"flag"

	"github.com/beshenkaD/unverblumt/bot"
)

func main() {
	version := "0.0.1"

	tokFlag := flag.String("token", "", "telegram bot token")
	debFlag := flag.Bool("debug", false, "use debug")

	flag.Parse()

	b := bot.New(*tokFlag, version, *debFlag)
	b.Run()
}
