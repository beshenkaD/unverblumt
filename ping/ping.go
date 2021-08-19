package ping

import "github.com/beshenkaD/unverblumt/bot"

func Ping(in *bot.CommandInput) (*bot.Output, error) {
	return &bot.Output{
		Text:    "pong",
		UseHTML: false,
		Photo:   nil,
	}, nil
}
