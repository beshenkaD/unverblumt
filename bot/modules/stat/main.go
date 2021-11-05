// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package stat

import (
	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	Module = &core.Module{
		Name:        "Statistics",
		Author:      "Beshenka",
		License:     "GNU GPL v2",
		Version:     "0.0.1",
		Description: "Provides basic bot functionality",

		ActiveCommands: map[string]core.ActiveCommand{
			"/stat": {
				Handler:     stat,
				Description: "some shit",
				Arguments:   []core.Argument{},
			},

			"/memory": {
				Handler:     memory,
				Description: "some shit",
				Arguments:   []core.Argument{},
			},
		},

		PollerFuncs: []func(*tb.Update) bool{
			collector,
		},
	}
)
