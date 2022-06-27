// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"github.com/beshenkaD/unverblumt/modules/me"
	"github.com/beshenkaD/unverblumt/modules/random"
	u "github.com/beshenkaD/unverblumt/unverblumt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	u.Run(
		&me.Me{},
		&random.Random{},
	)
}
