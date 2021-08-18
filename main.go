package main

import (
	"flag"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/hello"
	"github.com/beshenkaD/unverblumt/me"
)

func main() {
	version := "0.0.1"

	tokFlag := flag.String("token", "", "telegram bot token")
	debFlag := flag.Bool("debug", false, "use debug")

	flag.Parse()

	b := bot.New(*tokFlag, version, *debFlag)
	b.RegisterCommand("/me", "just a me command for telegram", me.Me)

	b.RegisterHook("greeter", "hello!!!", hello.Hello)

	b.Run()
}
