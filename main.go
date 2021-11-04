// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"github.com/beshenkaD/unverblumt/bot"
	"github.com/beshenkaD/unverblumt/internal/log"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var c bot.Config

	err := envconfig.Process("unverblumt", &c)
	if err != nil {
		log.Error.Fatal(err)
	}

	bot.Start(&c)
}
