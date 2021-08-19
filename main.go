package main

import (
	"flag"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/hello"
	"github.com/beshenkaD/unverblumt/me"
)

func main() {
	version := "0.1.3"

	tokFlag := flag.String("token", "", "telegram bot token")
	debFlag := flag.Bool("debug", false, "use debug")

	flag.Parse()

	b := bot.New(*tokFlag, version, *debFlag)
	b.RegisterCommand("/me", "Displays a message about yourself", me.Me)

	b.RegisterHook("greeter", "hello!!!", hello.Hello)

	b.Run()
}
