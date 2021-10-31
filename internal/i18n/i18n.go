// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package i18n

import (
	"io/ioutil"

	"github.com/leonelquinteros/gotext"
)

const (
	library         = "po"
	DefaultLanguage = "en_US"
)

var (
	available []string
)

func init() {
	t, _ := ioutil.ReadDir(library)

	for _, f := range t {
		available = append(available, f.Name())
	}

}

func getLocale(lang, domain string) *gotext.Locale {
	l := gotext.NewLocale(library, lang)
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
   Translate using specified domain
*/
func TD(l, d, val string, vars ...interface{}) string {
	return getLocale(l, d).Get(val, vars...)
}

/*
   Returns short version of language code
   Example: en_US -> en
*/
func Short(code string) string {
	return code[:2]
}

func GetAvailableLanguages() []string {
	return available
}
