// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/beshenkaD/unverblumt/core"
)

const (
	AppName = "UNVERBLUMT_"

	TokenEnv   = AppName + "TELEGRAM"
	ModulesEnv = AppName + "MODULES"
)

func getToken() string {
	t := os.Getenv(TokenEnv)

	if t == "" {
		log.Fatal(TokenEnv + " is unset!")
	}

	return t
}

func loadModules(u *core.Unverblumt) {
	raw := os.Getenv(ModulesEnv)

	if raw == "" {
		log.Println(ModulesEnv + " is unset. Loading without modules!")
		return
	}

	for _, path := range strings.Split(raw, " ") {
		m, err := u.LoadModule(path)

		if err != nil {
			fmt.Println(err)
			continue
		}

		u.RegisterModule(m)
	}
}
