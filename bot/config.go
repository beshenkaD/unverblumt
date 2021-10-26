// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/log"
	_ "github.com/joho/godotenv/autoload"
)

const (
	AppName = "UNVERBLUMT_"

	TokenEnv   = AppName + "TELEGRAM"
	ModulesEnv = AppName + "MODULES"
	TimeoutEnv = AppName + "TIMEOUT"
)

func getTimeout() time.Duration {
	t := os.Getenv(TimeoutEnv)

	if t == "" {
		log.Warn.Println(TimeoutEnv + " is unset. Using default value (60 seconds)")
		return 60 * time.Second
	}

	n, err := strconv.Atoi(t)
	if err != nil {
		log.Warn.Println(err)
		return 60 * time.Second
	}

	return time.Duration(n) * time.Second
}

func getToken() string {
	t := os.Getenv(TokenEnv)

	if t == "" {
		log.Error.Fatal(TokenEnv + " is unset!")
	}

	return t
}

func loadModules(u *core.Unverblumt) {
	raw := os.Getenv(ModulesEnv)

	if raw == "" {
		log.Warn.Println(ModulesEnv + " is unset. Loading without modules!")
		return
	}

	for _, path := range strings.Split(raw, " ") {
		m, err := u.LoadModule(path)

		if err != nil {
			log.Warn.Println(err)
			continue
		}

		u.RegisterModule(m)
	}
}
