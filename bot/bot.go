// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/log"
	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	base = &core.Module{
		Name:        "Main",
		Author:      "Beshenka",
		License:     "GNU GPL v2",
		Version:     "0.0.4",
		Description: "Provides basic bot functionality",
		ActiveCommands: map[string]core.ActiveCommand{
			"/help": {
				Handler:     help,
				Description: "sends you this message",
			},

			"/ping": {
				Handler:     ping,
				Description: "tests bot's reachability",
			},

			"/lang": {
				Handler:     lang,
				Description: "changes language for current chat",
			},

			"/start": {
				Handler:     start,
				Description: "sends you a welcome message",
			},
		},
	}
)

func Start() {
	u, err := core.New(tb.Settings{
		Token:     getToken(),
		Poller:    &tb.LongPoller{Timeout: getTimeout()},
		ParseMode: "HTML",
		OnError: func(err error, c tb.Context) {
			log.Warn.Println(err)
		},
	}, base)

	if err != nil {
		log.Error.Fatal(err)
	}

	loadModules(u)
	generateHelp(u)

	u.Start()
}
