// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package random

import (
	"math"
	"math/rand"
	"time"

	"github.com/beshenkaD/unverblumt/internal/log"
	u "github.com/beshenkaD/unverblumt/unverblumt"
	"gopkg.in/telebot.v3"
)

type Random struct{}

var myRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (m *Random) Init() error {
	log.Info.Println("initialized module `random`")
	return nil
}

func (m *Random) Commands() []*u.Command {
	return []*u.Command{
		{
			Cmd:     "/when",
			Desc:    "find out when the event will happen",
			Handler: when,
		},
	}
}

func when(c telebot.Context) error {
	var x = int(math.Round(myRand.ExpFloat64() * 3))

	t := time.Now().AddDate(x, myRand.Intn(12), myRand.Intn(31))

	return c.Reply(t.Format("02.01.2006"))
}
