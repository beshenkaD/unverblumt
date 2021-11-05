// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package settings

import (
	"sync"

	tb "gopkg.in/tucnak/telebot.v3"
)

/*
   Maybe use database?
*/

type lang struct {
	lock  sync.RWMutex
	langs map[int64]string
}

func (l *lang) Set(chatID int64, lang string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.langs[chatID] = lang
}

func (l *lang) Get(c tb.Context) string {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if lang, ok := l.langs[c.Chat().ID]; ok {
		return lang
	}

	return c.Sender().LanguageCode
}

func newLang() *lang {
	return &lang{
		lock:  sync.RWMutex{},
		langs: make(map[int64]string),
	}
}

var (
	Lang = newLang()
)
