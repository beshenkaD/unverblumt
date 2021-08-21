package bot

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sh "github.com/mattn/go-shellwords"
)

var (
	ErrNoArgs    = errors.New("no arguments passed")
	ErrWrongArgs = errors.New("wrong arguments passed")
)

type Bot struct {
	version   string
	startTime time.Time
	processed uint

	debug    bool
	Tg       *tgbotapi.BotAPI
	commands map[string]Command
	hooks    []Hook
}

func New(token, version string, debug bool) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug

	return &Bot{
		version:   version,
		processed: 0,
		debug:     debug,
		Tg:        bot,
		commands:  make(map[string]Command),
		hooks:     []Hook{},
	}
}

type Output struct {
	Text     string
	UseHTML  bool
	Photo    interface{}
	PhotoURL string
}

type commandFunc func(*CommandInput) (*Output, error)

type CommandInput struct {
	Command string
	Args    []string
	Msg     *tgbotapi.Message
	Bot     *Bot
}

type CommandParam struct {
	Name      string
	Desc      string
	Optional  bool
	Subparams []CommandParam
}

type Command struct {
	Name   string
	Desc   string
	Params []CommandParam
	Func   commandFunc
}

type hookFunc func(*HookInput) (*Output, error)

type HookInput struct {
	Msg *tgbotapi.Message
	Bot *Bot
}

type Hook struct {
	Name string
	Desc string
	Func hookFunc
}

func (b *Bot) RegisterCommand(name, desc string, params []CommandParam, f commandFunc) {
	b.commands[name] = Command{
		Name:   name,
		Desc:   desc,
		Params: params,
		Func:   f,
	}
}

func (b *Bot) RegisterHook(name, desc string, f hookFunc) {
	b.hooks = append(b.hooks, Hook{
		Name: name,
		Desc: desc,
		Func: f,
	})
}

func (b *Bot) sendText(msg string, useHTML bool, disablePreview bool, chat int64) {
	m := tgbotapi.NewMessage(chat, msg)
	m.DisableWebPagePreview = disablePreview

	if useHTML {
		m.ParseMode = tgbotapi.ModeHTML
	}

	_, err := b.Tg.Send(m)

	if err != nil {
		log.Println(err.Error())
	}
}

func (b *Bot) sendMessage(msg tgbotapi.Chattable) {
	_, err := b.Tg.Send(msg)

	if err != nil {
		log.Println(err.Error())
	}
}

func (b *Bot) sendError(err error, chat int64) {
	b.sendText("Error: "+err.Error(), false, true, chat)
}

func (b *Bot) parse(m *tgbotapi.Message) (interface{}, error) {
	if m.IsCommand() {
		args, err := sh.Parse(m.Text)

		p := "@" + b.Tg.Self.UserName
		if strings.Contains(args[0], p) {
			args[0] = args[0][0 : len(args[0])-len(p)]
		}

		if err != nil {
			return nil, err
		}

		return &CommandInput{
			Command: args[0],
			Args:    args[1:],
			Msg:     m,
			Bot:     b,
		}, nil
	}

	return &HookInput{
		Msg: m,
		Bot: b,
	}, nil
}

func (b *Bot) handleOut(out *Output, chat int64) {
	if out == nil {
		return
	}

	if out.Text != "" {
		b.sendText(out.Text, out.UseHTML, false, chat)
	}
	if out.PhotoURL != "" {
		msg := tgbotapi.NewPhotoUpload(chat, nil)
		msg.FileID = out.PhotoURL
		msg.UseExisting = true
		b.sendMessage(msg)
	}
	if out.Photo != nil {
		msg := tgbotapi.NewPhotoUpload(chat, out.Photo)
		b.sendMessage(msg)
	}
}

func (b *Bot) handleCommand(input *CommandInput) {
	if command, ok := b.commands[input.Command]; ok {
		out, err := command.Func(input)
		if err != nil {
			b.sendError(err, input.Msg.Chat.ID)
			return
		}

		b.handleOut(out, input.Msg.Chat.ID)
	}
}

func (b *Bot) handleHook(input *HookInput) {
	for _, hook := range b.hooks {
		out, err := hook.Func(input)
		if err != nil {
			b.sendError(err, input.Msg.Chat.ID)
			return
		}

		b.handleOut(out, input.Msg.Chat.ID)
	}
}

func (b *Bot) handleHelp(input *CommandInput) {
	if len(input.Args) > 0 {
		found := false
		for _, arg := range input.Args {
			if command, ok := b.commands[arg]; ok {
				found = true
				t := fmt.Sprintf("%s: %s.\n\n", command.Name, command.Desc)

				if command.Params != nil {
					t += "Available params:\n"
					for _, param := range command.Params {
						t = "\n" + t + fmt.Sprintf("%s\n    %s", param.Name, param.Desc)
						for _, subparam := range param.Subparams {
							t = "\n" + t + fmt.Sprintf("%s\n        %s", subparam.Name, subparam.Desc)
						}
						t += "\n"
					}
				}
				b.sendText(t, false, true, input.Msg.Chat.ID)
			}

			for _, hook := range b.hooks {
				if hook.Name == arg {
					found = true
					t := fmt.Sprintf("%s: %s.", hook.Name, hook.Desc)
					b.sendText(t, false, true, input.Msg.Chat.ID)
					break
				}
			}
		}

		if !found {
			b.sendError(ErrWrongArgs, input.Msg.Chat.ID)
		}
		return
	}

	commands := []string{}
	for name := range b.commands {
		commands = append(commands, name)
	}

	hooks := []string{}
	for _, hook := range b.hooks {
		hooks = append(hooks, hook.Name)
	}
	t := fmt.Sprintf("type /help <command>|<hook> to get detailed description\n\nAvailable commands: %v\nActive hooks: %v", strings.Join(commands, ", "), strings.Join(hooks, ", "))
	b.sendText(t, false, true, input.Msg.Chat.ID)
}

func (b *Bot) handleStat(input *CommandInput) {
	toMB := func(b uint64) uint64 {
		return b / 1024 / 1024
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	f := `%s v%s

Uptime: %s
Messages processed: %d
Memory usage: 
    Alloc: %v MB
    Sys: %v MB
    Heap (in use): %v MB
`
	t := fmt.Sprintf(f, b.Tg.Self.UserName, b.version, time.Since(b.startTime), b.processed, toMB(m.Alloc), toMB(m.Sys), toMB(m.HeapInuse))
	b.sendText(t, false, true, input.Msg.Chat.ID)
}

func (b *Bot) handleStart(input *CommandInput) {
	start := `Welcome to <b>%s v%s</b>!

This bot has a very simple and limited structure. It supports only <b><i>commands</i></b> and <b><i>hooks</i></b>. 
Commands are called using the prefix. Example:
    <code>/ping</code>
or using the prefix with the name of the bot:
    <code>/ping@%s</code>
Hooks work automatically if bot has access to messages.

Commands can be called with arguments. For example: 
    <code>/weather Moscow</code> 
If the argument contains several words, then use quotes:
    <code>/weather "Saint Petersburg"</code>

Source code is available here: github.com/beshenkaD/unverblumt

Use <b>/help</b> for more information. 
`
	start = fmt.Sprintf(start, b.Tg.Self.UserName, b.version, b.Tg.Self.UserName)
	b.sendText(start, true, true, input.Msg.Chat.ID)
}

func (b *Bot) messageReceived(msg *tgbotapi.Message) {
	b.processed++

	a, err := b.parse(msg)

	if err != nil {
		b.sendError(err, msg.Chat.ID)
	}

	switch a := a.(type) {
	case *CommandInput:
		if a.Command == "/help" {
			b.handleHelp(a)
			return
		}
		if a.Command == "/stat" {
			b.handleStat(a)
			return
		}
		if a.Command == "/start" {
			b.handleStart(a)
			return
		}

		go b.handleCommand(a)
	case *HookInput:
		go b.handleHook(a)
	}
}

func (b *Bot) Run() {
	b.startTime = time.Now()

	log.Printf("Authorized on account %s", b.Tg.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.Tg.GetUpdatesChan(u)
	if err != nil {
		log.Println(err.Error())
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.messageReceived(update.Message)
	}
}

func (b *Bot) PrintCommands() {
	fmt.Println("Pass this strings to BotFather")
	fmt.Println("")
	for name, command := range b.commands {
		fmt.Printf("%s - %s\n", name[1:], command.Desc)
	}

	fmt.Println("start - start")
	fmt.Println("help - displays help")
	fmt.Println("stat - displays bot statistics")
}
