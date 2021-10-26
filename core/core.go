// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package core

import (
	"plugin"

	"github.com/beshenkaD/unverblumt/internal/log"

	tb "gopkg.in/tucnak/telebot.v3"
)

/*
   Represents any command that can be passed to bot
*/
type Command struct {
	Handler     tb.HandlerFunc
	Arguments   string
	Description string
}

/*
   Module unites active and passive commands in logical blocks.
   ActiveCommands  - commands called with '/' symbol. Example: `/ping`
   PassiveCommands - commands called on some event. Example: `tb.OnText`
*/
type Module struct {
	Name            string
	Author          string
	License         string
	Version         string
	Description     string
	ActiveCommands  map[string]Command
	PassiveCommands map[string][]Command
}

/*
   Main bot object
*/
type Unverblumt struct {
	Bot *tb.Bot

	Modules map[string]*Module
}

/*
   Adds a module with the given name to the bot.
*/
func (u *Unverblumt) RegisterModule(m *Module) {
	name := m.Name

	u.Modules[name] = m

	for k, v := range m.ActiveCommands {
		u.Bot.Handle(k, v.Handler)
	}

	for k, v := range m.PassiveCommands {
		m.PassiveCommands[k] = v
		u.setHandler(k)
	}

	log.Info.Printf("Module `%s` has been registered", name)
}

/*
   A little hack for having multiple handlers for same event type.
*/
func (u *Unverblumt) setHandler(event string) {
	u.Bot.Handle(event, func(ctx tb.Context) error {
		for _, m := range u.Modules {
			for _, c := range m.PassiveCommands[event] {
				c.Handler(ctx)
			}
		}
		return nil
	})
}

/*
   Loads the shared library (go plugin) containing the module.
   Your plugin must export function `Init() *core.Module`
   It is called to extract module.
*/
func (u *Unverblumt) LoadModule(path string) (*Module, error) {
	p, err := plugin.Open(path)

	if err != nil {
		return nil, err
	}

	f, err := p.Lookup("Init")

	if err != nil {
		return nil, err
	}

	return f.(func() *Module)(), nil
}

/*
   Creates new bot instance.
*/
func New(token string, parseMode string, verbose bool) (*Unverblumt, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:     token,
		Poller:    &tb.LongPoller{Timeout: 60 * time.Second},
		Verbose:   verbose,
		ParseMode: parseMode,
	})

	if err != nil {
		return nil, err
	}

	u := &Unverblumt{
		Bot:     bot,
		Modules: make(map[string]*Module),
	}

	u.registerBase()

	return u, nil
}

func (u *Unverblumt) Start() {
	u.generateHelpData()

	u.Bot.Start()
}
