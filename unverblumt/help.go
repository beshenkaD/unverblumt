// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package unverblumt

import (
	"fmt"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"github.com/beshenkaD/unverblumt/internal/style"
)

func (u *unverblumt) GenerateHelp(lang string, args []string) string {
	if len(args) == 0 {
		return i18n.T(lang, "preamble_") + "\n" + u._brief(lang)
	}

	if c, ok := u.commands[args[0]]; ok {
		return _detail(c, lang)
	}

	return "fd"
}

func (u *unverblumt) _brief(lang string) string {
	var s = ""

	for _, c := range u.commands {
		s += c.Cmd + "  :  " + i18n.TD(lang, "help", c.Desc) + "\n"
	}

	return s
}

func _detail(c *Command, lang string) string {
	arg := func(a Arg) string {
		return fmt.Sprintf("%s : %s", i18n.TD(lang, "help", a.Name), i18n.TD(lang, "help", a.Desc))
	}

	enquote := func(s, o, c string) string {
		return o + s + c
	}

	var required = ""
	var optional = ""
	var descriptions = ""

	for _, a := range c.Args {
		n := i18n.TD(lang, "help", a.Name)

		if a.Required {
			required += n + " "
		} else {
			optional += n + " "
		}

		descriptions += arg(a) + "\n    "
	}

	if required != "" {
		required = enquote(required, "{ ", "}")
		required = style.Underline(required)
	}

	if optional != "" {
		optional = enquote(optional, "[ ", "]")
		optional = style.Bold(optional)
	}

	return fmt.Sprintf("%s %s %s : %s\n    %s", c.Cmd, required, optional, i18n.TD(lang, "help", c.Desc), descriptions)
}
