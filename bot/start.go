// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"github.com/beshenkaD/unverblumt/bot/settings"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

func start(c tb.Context) error {
	if c.Message().Private() {
		msg := i18n.T(
			settings.Lang.Get(c.Chat().ID),
			"not implemented",
		)

		return c.Send(msg, &tb.SendOptions{
			DisableWebPagePreview:   true,
			DisableContentDetection: true,
		})
	}

	return nil
}
