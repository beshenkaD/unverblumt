// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package core

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v3"
)

/*
   I think this should be refactored. This code looks like shit
   Separate help generator from base?
   Create all data at startup then use it without bothering main object?
   Use FSM for buttons?
   Use templates for messages?
*/
var (
	// [Modules] [Commands] [About]
	baseSelector = &tb.ReplyMarkup{}

	m tb.Btn // `Modules` button
	c tb.Btn // `Commands` button

	// [/command1]
	// [/command2]
	// ...
	commandsSelector = &tb.ReplyMarkup{}
	_commands        = make(map[string]string) // Menu item text by ID

	// [module1]
	// [module2]
	// ...
	modulesSelector = &tb.ReplyMarkup{}
	_modules        = make(map[string]string)
)

/*
   Returns pretty-printed information about module
*/
func getModuleInfo(m *Module) string {
	const s = `<b>%s</b> module
    %s.

<i>    <b>Author:</b> %s
    <b>License:</b> %s
    <b>Version:</b> %s</i>

<b>Active commands</b>:
    <b>%s</b>

<b>Passive commands</b>:
    <b>%s</b>`

	// Rewrite with templates?
	var (
		act string
		pas string
	)

	if len(m.ActiveCommands) == 0 {
		act = "none"
	} else {
		cmds := []string{}

		for n, c := range m.ActiveCommands {
			cmds = append(cmds, n+"</b> — "+c.Description+"<b>")
		}

		act = strings.Join(cmds, ",\n    ")
	}

	if len(m.PassiveCommands) == 0 {
		pas = "none"
	} else {
		cmsd := []string{}

		for _, c := range m.PassiveCommands {
			descs := []string{}

			for _, h := range c {
				descs = append(descs, h.Description)
			}

			cmsd = append(cmsd, strings.Join(descs, ",\n    "))
		}

		pas = strings.Join(cmsd, ",\n    ")
	}

	return fmt.Sprintf(s, m.Name, m.Description, m.Author, m.License, m.Version, act, pas)
}

/*
   This function generates "modules" menu inside help message
*/
func (u Unverblumt) generateModuleButtons() {
	kb := [][]tb.InlineButton{}

	for name, module := range u.Modules {
		id := uuid.NewString()

		b := modulesSelector.Data(name, id).Inline()
		_modules[id] = getModuleInfo(module)

		u.Bot.Handle(b, func(c tb.Context) error {
			c.Respond(&tb.CallbackResponse{})
			return c.Edit(_modules[b.Unique])
		})

		kb = append(kb, []tb.InlineButton{*b})
	}

	modulesSelector.InlineKeyboard = kb
}

/*
   Returns pretty-printed information about command
*/
func commandInfo(name string, c Command) string {
	const s = `<b>Usage:</b> <b>%s</b> <i>%s</i>

    %s`
	return fmt.Sprintf(s, name, c.Arguments, c.Description)
}

/*
   This function generates "commands" menu inside help message
*/
func (u Unverblumt) generateCommandButtons() {
	kb := [][]tb.InlineButton{}

	for _, module := range u.Modules {
		for name, command := range module.ActiveCommands {
			id := uuid.NewString()

			b := commandsSelector.Data(name, id).Inline()
			_commands[id] = commandInfo(name, command)

			u.Bot.Handle(b, func(c tb.Context) error {
				c.Respond(&tb.CallbackResponse{})
				return c.Edit(_commands[b.Unique])
			})

			kb = append(kb, []tb.InlineButton{*b})
		}
	}
	commandsSelector.InlineKeyboard = kb
}

/*
   This function generates help data message using modules
   so it should be called after loading all modules
*/
func (u Unverblumt) generateHelpData() {
	m = baseSelector.Data("Modules", "m")
	c = baseSelector.Data("Commands", "c")

	baseSelector.Inline(
		baseSelector.Row(m, c),
	)

	u.generateModuleButtons()
	u.generateCommandButtons()

	u.Bot.Handle(&m, func(c tb.Context) error {
		c.Respond(&tb.CallbackResponse{})
		return c.Edit("<b>Choose module:</b>", modulesSelector)
	})

	u.Bot.Handle(&c, func(c tb.Context) error {
		c.Respond(&tb.CallbackResponse{})
		return c.Edit("<b>Choose command:</b>", commandsSelector)
	})
}

/*
   This wrapper is needed to access Unverblumt object from command
   Maybe store needed information somewhere around?
*/
func (u Unverblumt) generateHelp() tb.HandlerFunc {
	help := func(c tb.Context) error {
		if len(c.Args()) > 0 {
			for _, arg := range c.Args() {
				if m, ok := u.Modules[arg]; ok {
					c.Send(getModuleInfo(m))
				}
			}

			return nil
		}

		const s = `You can click on buttons, or use the following syntax:
    <b>/help [command] | [module]</b>`
		return c.Send(s, baseSelector)
	}

	return help
}
