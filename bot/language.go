// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"strings"

	st "github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

func language(c tb.Context) error {
	available := i18n.GetAvailableLanguages()
	current := st.Lang.Get(c.Chat().ID)

	if len(c.Args()) == 0 {
		return c.Send(i18n.T(current, "Available languages: %v",
			strings.Join(available, ", ")))
	}

	in := c.Args()[0]

	for _, a := range i18n.GetAvailableLanguages() {
		if in == a {
			st.Lang.Set(c.Chat().ID, in)
			return c.Send(i18n.T(in, "Language `%s` successfully installed", in))
		}
	}

	return c.Send(i18n.T(current, "Invalid value: `%s`", in))
}
