package me

import (
	"strings"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"github.com/beshenkaD/unverblumt/internal/log"
	"github.com/beshenkaD/unverblumt/internal/style"
	u "github.com/beshenkaD/unverblumt/unverblumt"
	"gopkg.in/telebot.v3"
)

type Me struct{}

func (m *Me) Init() error {
	log.Info.Println("initialized module `me`")
	return nil
}

func (m *Me) Commands() []*u.Command {
	return []*u.Command{
		{
			Cmd:  "/me",
			Desc: "command from IRC",
			Args: []u.Arg{
				{
					Name:     "action",
					Desc:     "the action you want to perform",
					Required: true,
				},
			},
			Handler: me,
		},
	}
}

func me(c telebot.Context) error {
	if len(c.Args()) == 0 {
		return c.Reply(i18n.T(c.Sender().LanguageCode, "please provide an action"))
	}

	return c.Send(
		style.Bold(c.Sender().FirstName + " " + c.Sender().LastName + " " + strings.Join(c.Args(), " ")),
	)
}
