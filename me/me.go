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
	if l != "" {
		f += " "
	}

	t := fmt.Sprintf("<b>%s%s %s</b>", f, l, strings.Join(in.Args, " "))

	in.Bot.DeleteMessage(in.Msg.Chat.ID, in.Msg.MessageID)

	return &bot.Output{
		Text:    t,
		UseHTML: true,
		Photo:   nil,
	}, nil
}
