// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)

const s = `Welcome to <b>Unverblümt</b>!

<b>Unverblümt</b> is a general purpose telegram bot. It supports active and passive commands (hooks), inline buttons and inline queries.

Source code is available here: github.com/beshenkaD/unverblumt
Use <b>/help</b> for more information about usage.`

func start(c tb.Context) error {
	if c.Message().Private() {
		return c.Send(s, &tb.SendOptions{
			DisableWebPagePreview:   true,
			DisableContentDetection: true,
		})
	}

	return nil
}

func Init() *core.Module {
	return &core.Module{
		Name:        "Start",
		License:     "GNU GPL v2",
		Author:      "Beshenka",
		Version:     "0.0.1",
		Description: "Defines a `/start` command",

		ActiveCommands: map[string]core.Command{
			"/start": {
				Handler:     start,
				Description: "Welcome message. Contains brief overview of bot",
			},
		},
	}
}
