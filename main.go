// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"log"

	"github.com/beshenkaD/unverblumt/core"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	u, err := core.New(getToken(), "HTML", false)

	if err != nil {
		log.Fatal(err)
	}

	loadModules(u)
	u.Start()
}
