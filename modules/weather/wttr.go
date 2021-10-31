// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"net/url"

	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	wttr = core.ActiveCommand{
		Handler:     _wttr,
		Description: "gets information from wttr.in",
	}
)

func _wttr(c tb.Context) error {
	vals := url.Values{}
	vals.Add("p", "")
	vals.Add("0", "")
	vals.Add("lang", "en")

	var (
		city = "Moscow.png"
	)

	base, _ := url.Parse("https://www.wttr.in")

	base.Path += city
	base.RawQuery = vals.Encode()

	return c.Send(&tb.Photo{File: tb.FromURL(base.String())})
}
