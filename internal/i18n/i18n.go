// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package i18n

import (
	"io/ioutil"

	"github.com/leonelquinteros/gotext"
)

const (
	DefaultLanguage = "en_US"
)

var (
	avail []string
)

func init() {
	gotext.Configure("po", DefaultLanguage, "default")

	t, _ := ioutil.ReadDir(gotext.GetLibrary())

	for _, f := range t {
		avail = append(avail, f.Name())
	}

}

func T(lang string, val string, vars ...interface{}) string {
	// There is a mutex inside, don't worry
	gotext.SetLanguage(lang)

	return gotext.Get(val, vars...)
}

func GetAvailableLanguages() []string {
	return avail
}
