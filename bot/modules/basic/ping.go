// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package basic

import (
	st "github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

func ping(c tb.Context) error {
	return c.Send(i18n.T(st.Lang.Get(c.Chat().ID), "pong"))
}
