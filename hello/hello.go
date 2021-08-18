package hello

import (
	"strings"

	"github.com/beshenkaD/unverblumt/bot"
)

func Hello(in *bot.HookInput) (*bot.Output, error) {
	if strings.Contains(in.Msg.Text, "hello") {
		return &bot.Output{
			Text:      "hello!",
			PhotoPath: "",
		}, nil
	}

	return nil, nil
}
