// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"strings"
	"time"

	"github.com/beshenkaD/unverblumt/bot/modules/basic"
	"github.com/beshenkaD/unverblumt/bot/modules/stat"
	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/log"
	tb "gopkg.in/tucnak/telebot.v3"
)

type Config struct {
	Telegram string
	Modules  []string
	Timeout  time.Duration
}

func load(u *core.Unverblumt, modules []string) {
	if len(modules) == 0 {
		log.Warn.Println("Loading without external modules")
		return
	}

	for _, path := range modules {
		path = strings.TrimSpace(path) + ".so"
		m, err := u.LoadModule(path)

		if err != nil {
			log.Warn.Println(err)
			continue
		}

		u.RegisterModule(m)
	}
}

func Start(c *Config) {
	u, err := core.New(tb.Settings{
		Token:     c.Telegram,
		Poller:    &tb.LongPoller{Timeout: c.Timeout},
		ParseMode: "HTML",
		OnError: func(err error, c tb.Context) {
			log.Warn.Println(err)
		},
	}, basic.Module, stat.Module)

	if err != nil {
		log.Error.Fatal(err)
	}

	load(u, c.Modules)

	basic.GenerateHelp(u)

	u.Start()
}
