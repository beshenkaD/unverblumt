// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package unverblumt

import (
	"github.com/beshenkaD/unverblumt/internal/i18n"
	"gopkg.in/telebot.v3"
)

func loadMiddlewares() {
	ignoreFwdCommands()
}

func loadBuiltins() {
	ping()
	help()
	start()
}

func ping() {
	Get().AddCommand(&Command{
		Cmd:  "/ping",
		Desc: "check if bot is active",
		Args: nil,
		Handler: func(c telebot.Context) error {
			return c.Reply(i18n.T(c.Message().Sender.LanguageCode, "pong"))
		},
	},
	)
}

func help() {
	Get().AddCommand(&Command{
		Cmd:  "/help",
		Desc: "usage for the bot",
		Args: []Arg{
			{
				Name:     "name",
				Desc:     "show usage for given command name",
				Required: false,
			},
		},
		Handler: func(c telebot.Context) error {
			return c.Reply(Get().GenerateHelp(c.Sender().LanguageCode, c.Args()))
		},
	})
}

func start() {
	Get().AddCommand(&Command{
		Cmd:  "/start",
		Desc: "start message",
		Args: nil,
		Handler: func(c telebot.Context) error {
			if c.Message().Private() {
				msg := i18n.T(c.Sender().LanguageCode, "start_")

				return c.Send(msg, &telebot.SendOptions{
					DisableWebPagePreview: true,
				})
			}
			return nil
		},
	})
}

func ignoreFwdCommands() {
	Get().Bot.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			if msg := c.Message(); msg != nil && msg.IsForwarded() {
				return nil
			}

			return next(c)
		}
	})
}
