package me

import (
	"fmt"
	"strings"

	"github.com/beshenkaD/unverblumt/bot"
)

func Me(in *bot.CommandInput) (*bot.Output, error) {
	if len(in.Args) < 1 {
		return nil, bot.ErrNoArgs
	}

	f := strings.Title(in.Msg.From.FirstName)
	l := strings.Title(in.Msg.From.LastName)
	t := fmt.Sprintf("_%s %s %s_", f, l, strings.Join(in.Args, " "))

	return &bot.Output{
		Text:      t,
		PhotoPath: "",
	}, nil
}
