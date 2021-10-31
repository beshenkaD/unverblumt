// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"strings"

	st "github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	tb "gopkg.in/tucnak/telebot.v3"
)

func add(c tb.Context, btn *tb.Btn, code, full string) {
	c.Bot().Handle(btn, func(ct tb.Context) error {
		st.Lang.Set(c.Chat().ID, code)

		c.Send(i18n.T(code, "Now you are using %s", full))

		return ct.Respond(&tb.CallbackResponse{})
	})
}

func lang(c tb.Context) error {
	current := st.Lang.Get(c.Chat().ID)

	selector := &tb.ReplyMarkup{}

	buttons := []tb.Btn{}

	for _, lang := range i18n.GetAvailableLanguages() {
		full := strings.Title(display.Self.Name(language.Make(lang)))
		btn := selector.Data(full, lang)

		add(c, &btn, lang, full)

		buttons = append(buttons, btn)
	}

	kb := [][]tb.InlineButton{}
	for _, btn := range buttons {
		kb = append(kb, []tb.InlineButton{*btn.Inline()})
	}

	selector.InlineKeyboard = kb

	return c.Send(i18n.T(current, "Choose your language"), selector)
}
