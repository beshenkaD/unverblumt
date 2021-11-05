// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package basic

import (
	"github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

const startMessage = `<b>Welcome!</b>
Unverblümt is a general purpose telegram bot designed to be comfortable and powerful.

Source code is available here: github.com/beshenkaD/unverblumt
Use <b>/help</b> to get information about usage.`

func start(c tb.Context) error {
	if c.Message().Private() {
		msg := i18n.T(
			settings.Lang.Get(c),
			startMessage,
		)

		return c.Send(msg, &tb.SendOptions{
			DisableWebPagePreview:   true,
			DisableContentDetection: true,
		})
	}

	return nil
}
