// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package text

import (
	"github.com/beshenkaD/unverblumt/internal/log"
	u "github.com/beshenkaD/unverblumt/unverblumt"
)

type Text struct{}

func (m *Text) Init() error {
	log.Info.Println("initialized module `text`")
	return nil
}

func (m *Text) Commands() []*u.Command {
	return []*u.Command{
		{
			Cmd:  "/flip",
			Desc: "flips string upside-down",
			Args: []u.Arg{
				{
					Name:     "string",
					Desc:     "string to flip",
					Required: true,
				},
			},
			Handler: flip,
		},
	}
}
