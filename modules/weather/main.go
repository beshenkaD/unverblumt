// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"github.com/beshenkaD/unverblumt/core"
)

func Init() *core.Module {
	return &core.Module{
		Name:        "Weather",
		Author:      "Beshenka",
		License:     "GNU GPL v2",
		Version:     "0.0.1",
		Description: "ummm weather i guess",

		ActiveCommands: map[string]core.ActiveCommand{
			"/wttr": wttr,
		},
	}
}
