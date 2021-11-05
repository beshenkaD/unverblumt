// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"net/http"
	"net/url"

	"github.com/beshenkaD/unverblumt/bot/args"
	"github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	wttr = core.ActiveCommand{
		Handler:     _wttr,
		Description: "gets current weather from wttr.in",

		Arguments: []core.Argument{
			{Name: "-loc", Description: "location for which the weather is requested. It might be: city name in any language, any location (example: Eiffel tower), airport code, area code or gps coordinates. Defaults to \"Moscow\"", Required: false},
			{Name: "-lang", Description: "the language in which the weather forecast will be. It must be 2-letter code like: en, ru, fr etc. See `https://wttr.in/:help` for list of supported languages. Defaults to language set with /lang command", Required: false},
		},
	}
)

func _wttr(c tb.Context) error {
	base, _ := url.Parse("https://www.wttr.in")

	var a struct {
		Loc  string
		Lang string
	}
	a.Loc = "Moscow"
	a.Lang = settings.Lang.Get(c)

	if err := args.Parse(c.Text(), &a); err != nil {
		return c.Send(i18n.T(a.Lang, err.Error()))
	}

	vals := url.Values{}
	vals.Add("p", "")
	vals.Add("0", "")
	vals.Add("lang", a.Lang)

	base.Path += a.Loc + ".png"
	base.RawQuery = vals.Encode()

	r, _ := http.Get(base.String())

	return c.Send(&tb.Photo{File: tb.FromReader(r.Body)})
}
