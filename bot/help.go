// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)

/*
   Create all data at startup then use it without bothering main object
   Use FSM for buttons
   Use templates for messages
*/

/*
   This function should pre-allocate all help information
*/
func generateHelp(u *core.Unverblumt) {
}

func help(c tb.Context) error {
	return c.Send("Not implemented yet")
}
