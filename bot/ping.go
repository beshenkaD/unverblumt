// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	tb "gopkg.in/tucnak/telebot.v3"
)

func ping(c tb.Context) error {
	return c.Send("pong")
}
