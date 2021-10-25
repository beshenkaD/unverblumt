// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package core

import (
	tb "gopkg.in/tucnak/telebot.v3"
)

func ping(c tb.Context) error {
	c.Send("pong")

	return nil
}

func (u *Unverblumt) registerBase() {
	u.RegisterModule(&Module{
		Name:        "Base",
		License:     "GNU GPL v2",
		Author:      "Beshenka",
		Version:     "0.0.1",
		Description: "Contains very minimal bot functionality",

		ActiveCommands: map[string]Command{
			"/ping": {
				Handler:     ping,
				Description: "A simple way to check bot's availability",
			},

			"/help": {
				Handler:     u.generateHelp(),
				Description: "Displays help message",
			},
		},
	})
}
