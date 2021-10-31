// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

/*
   This package is responsive for low-level bot organization
   All it's point is managing modules.
*/
package core

import (
	"plugin"

	"github.com/beshenkaD/unverblumt/internal/log"

	tb "gopkg.in/tucnak/telebot.v3"
)

type Argument struct {
	Name        string
	Description string
	Required    bool
}

type ActiveCommand struct {
	Handler     tb.HandlerFunc
	Description string
	Arguments   []Argument
}

type PassiveCommand struct {
	Handler     tb.HandlerFunc
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
	ActiveCommands  map[string]ActiveCommand
	PassiveCommands map[string][]PassiveCommand
}

/*
   Main bot object
*/
type Unverblumt struct {
	Bot *tb.Bot

	modules map[string]*Module
}

/*
   Returns copy of all registered modules
   This should be useful for various info commands,
   and mainly for the 'help' command
*/
func (u *Unverblumt) GetModules() map[string]Module {
	r := make(map[string]Module)

	for n, m := range u.modules {
		r[n] = *m
	}

	return r
}

/*
   Adds a module with the given name to the bot.

   maybe restrict registering modules at runtime?
*/
func (u *Unverblumt) RegisterModule(m *Module) {
	name := m.Name

	u.modules[name] = m

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
		for _, m := range u.modules {
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
func New(s tb.Settings, m ...*Module) (*Unverblumt, error) {
	bot, err := tb.NewBot(s)

	if err != nil {
		return nil, err
	}

	u := &Unverblumt{
		Bot:     bot,
		modules: make(map[string]*Module),
	}

	for _, module := range m {
		u.RegisterModule(module)
	}

	return u, nil
}

func (u *Unverblumt) Start() {
	log.Info.Println("Bot has been started with the name: " + u.Bot.Me.FirstName)

	u.Bot.Start()
}
