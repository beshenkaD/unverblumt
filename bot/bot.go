// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/log"
	tb "gopkg.in/tucnak/telebot.v3"
)

func Start() {
	base := &core.Module{
		Name:        "a",
		Author:      "",
		License:     "",
		Version:     "",
		Description: "",
		ActiveCommands: map[string]core.Command{
			"/help": {
				Handler:     generateHelp(),
				Arguments:   "",
				Description: "",
			},

			"/ping": {
				Handler: func(c tb.Context) error {
					return c.Send("pong")
				},
				Arguments:   "",
				Description: "",
			},
		},
	}

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
	generateHelpData(u, u.GetModules())

	u.Start()
}
