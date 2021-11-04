// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package basic

import "github.com/beshenkaD/unverblumt/core"

var (
	Module = &core.Module{
		Name:        "Basic",
		Author:      "Beshenka",
		License:     "GNU GPL v2",
		Version:     "0.0.4",
		Description: "Provides basic bot functionality",

		ActiveCommands: map[string]core.ActiveCommand{
			"/help": {
				Handler:     help,
				Description: "sends you this message",
				Arguments: []core.Argument{
					{Name: "-c", Description: "command about which you need to get help", Required: false},
					{Name: "-m", Description: "module about which you need to get help", Required: false},
				},
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
