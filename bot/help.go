// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package bot

import (
	"fmt"
	"sync"

	"github.com/beshenkaD/unverblumt/core"
	"github.com/beshenkaD/unverblumt/internal/i18n"
	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	m        = sync.Mutex{}
	commands = make(map[string]core.ActiveCommand)
	modules  = make(map[string]core.Module)
)

func generateHelp(u *core.Unverblumt) {
	for n, m := range u.GetModules() {
		modules[n] = m
	}

	for _, m := range modules {
		for n, c := range m.ActiveCommands {
			commands[n] = c
		}
	}
}

/*
   <Brief overview>

   Module name
       /command1 - description
       /command2 - description

   Module name
       /command1 - description
       /command2 - description
   ...
*/
func getJustHelp(lang string) string {
	f := "%s\n\n"

	m.Lock()
	defer m.Unlock()

	for n, m := range modules {
		name := "<b>" + i18n.T(lang, n) + "</b>"

		mf := name + "\n"

		for t, c := range m.ActiveCommands {
			mf += fmt.Sprintf("    <b>%s</b> — %s\n", t, i18n.T(lang, c.Description))
		}

		f += mf + "\n"
	}

	return fmt.Sprintf(f, i18n.T(lang, "brief overview"))
}

/*
   <command>: a command that do something cool

   Usage: <command> [options]

   Options:
        somearg1 - argument that defines something (required)
	    somearg2 - argument that defines something
*/
func getHelpForCommand(lang string, cmd string) (string, bool) {
	m.Lock()
	defer m.Unlock()

	r := func(b bool) string {
		if b {
			return "(" + i18n.T(lang, "required") + ")"
		}
		return ""
	}

	if c, ok := commands[cmd]; ok {
		f := "<b>%s</b>: %s\n\n"

		long := ""

		for _, arg := range c.Arguments {
			long += fmt.Sprintf(
				"    %s — %s <b>%s</b>\n",
				i18n.T(lang, arg.Name),
				i18n.T(lang, arg.Description),
				r(arg.Required),
			)
		}

		f += fmt.Sprintf(
			"<b>%s</b>: %s [%s]\n\n", i18n.T(lang, "Usage"), cmd, i18n.T(lang, "options"))

		if long != "" {
			f += "<b>" + i18n.T(lang, "Options") + "</b>" + ":\n"
			f += long
		}

		return fmt.Sprintf(f, cmd, i18n.T(lang, c.Description)), true
	}

	return "", false
}

/*
  <module>: a module that combines commands for something

  Info:
       Version: 0.0.1
       Author: Beshenka
       License: GNU GPL v2

   Active commands:
       /command1 - cool command. My favorite
       ...

   Passive commands:
       someFilter - filters something
       ...
*/
func getHelpForModule(lang string, mod string) (string, bool) {
	m.Lock()
	defer m.Unlock()

	if m, ok := modules[mod]; ok {
		f := "<b>%s</b>: %s\n\n"

		I := "<b>" + i18n.T(lang, "Info") + "</b>"
		V := i18n.T(lang, "Version")
		A := i18n.T(lang, "Author")
		L := i18n.T(lang, "License")

		f += fmt.Sprintf(`%s:
    %s: <i>%s</i>
    %s: <i>%s</i>
    %s: <i>%s</i>

`, I, V, m.Version, A, m.Author, L, m.License)

		active := ""
		for n, c := range m.ActiveCommands {
			active += "    " + n + " — " + i18n.T(lang, c.Description) + "\n"
		}

		if active != "" {
			f += "<b>" + i18n.T(lang, "Active commands") + "</b>" + ":\n"
			f += active
			f += "\n"
		}

		passive := ""
		for _, cmds := range m.PassiveCommands {
			for _, c := range cmds {
				passive += "    " + i18n.T(lang, c.Description) + "\n"
			}
		}

		if passive != "" {
			f += "<b>" + i18n.T(lang, "Passive commands") + "</b>" + ":\n"
			f += passive
		}

		return fmt.Sprintf(f, i18n.T(lang, m.Name), i18n.T(lang, m.Description)), true
	}

	return "", false
}

func help(c tb.Context) error {
	return c.Send("not implemented yet")
}
