// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"strings"

	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)

func me(c tb.Context) error {
	if len(c.Args()) == 0 {
		return nil
	}

	action := strings.Join(c.Args(), " ")

	return c.Send("<b>" + c.Sender().FirstName + " " + action + "</b>")
}

func Init() *core.Module {
	return &core.Module{
		Name:        "Me",
		Author:      "Beshenka",
		License:     "GNU GPL v2",
		Version:     "0.0.1",
		Description: "Provides /me command like in IRC",

		ActiveCommands: map[string]core.ActiveCommand{
			"/me": {
				Handler: me,
				Arguments: []core.Argument{
					{Name: "action", Description: "action you would perform", Required: true},
				},
				Description: "/me command from IRC",
			},
		},
	}
}
