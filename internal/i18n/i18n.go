// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package i18n

import (
	"github.com/leonelquinteros/gotext"
	"gopkg.in/telebot.v3"
)

func getLocale(lang, domain string) *gotext.Locale {
	l := gotext.NewLocale("po", lang)
	l.AddDomain(domain)

	return l
}

/*
   Translate using default domain
*/
func T(l, val string, vars ...interface{}) string {
	return getLocale(l, "default").Get(val, vars...)
}

/*
   Translate using default domain and language code from telegram
*/
func TC(c telebot.Context, val string, vars ...interface{}) string {
	return getLocale(c.Sender().LanguageCode, "default").Get(val, vars...)
}

/*
   Translate using given domain
*/
func TD(l, domain, val string, vars ...interface{}) string {
	return getLocale(l, domain).Get(val, vars...)
}
