package bot

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sh "github.com/mattn/go-shellwords"
)

var (
	ErrNoArgs = errors.New("no arguments passed")
)

type Bot struct {
	version   string
	startTime time.Time
	processed uint

	debug    bool
	tg       *tgbotapi.BotAPI
	commands map[string]Command
	hooks    []Hook
}

func New(token, version string, debug bool) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug

	return &Bot{
		version:   version,
		processed: 0,
		debug:     debug,
		tg:        bot,
		commands:  make(map[string]Command),
		hooks:     []Hook{},
	}
}

type commandFunc func(*CommandInput) (*Output, error)

type Output struct {
	Text      string
	PhotoPath string
}

type CommandInput struct {
	Command string
	Args    []string
	Msg     *tgbotapi.Message
	Bot     *Bot
}

type Command struct {
	Name string
	Desc string
	Func commandFunc
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

func (b *Bot) RegisterCommand(name, desc string, f commandFunc) {
	b.commands[name] = Command{
		Name: name,
		Desc: desc,
		Func: f,
	}
}

func (b *Bot) RegisterHook(name, desc string, f hookFunc) {
	b.hooks = append(b.hooks, Hook{
		Name: name,
		Desc: desc,
		Func: f,
	})
}

func (b *Bot) sendText(msg string, chat int64) {
	m := tgbotapi.NewMessage(chat, msg)
	m.ParseMode = tgbotapi.ModeMarkdown

	_, err := b.tg.Send(m)

	if err != nil {
		log.Println(err.Error())
	}
}

func (b *Bot) sendMessage(msg tgbotapi.Chattable) {
	_, err := b.tg.Send(msg)

	if err != nil {
		log.Println(err.Error())
	}
}

func (b *Bot) sendError(err error, chat int64) {
	b.sendText("Error: "+err.Error(), chat)
}

func (b *Bot) parse(m *tgbotapi.Message) (interface{}, error) {
	if m.IsCommand() {
		args, err := sh.Parse(m.Text)

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
		b.sendText(out.Text, chat)
	}
	if out.PhotoPath != "" {
		bytes, err := os.ReadFile(out.PhotoPath)
		if err != nil {
			log.Println(err.Error())
			return
		}

		msg := tgbotapi.NewPhotoUpload(chat, bytes)
		b.tg.Send(msg)
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
		for _, arg := range input.Args {
			if command, ok := b.commands[arg]; ok {
				t := fmt.Sprintf("%s: %s.", command.Name, command.Desc)
				b.sendText(t, input.Msg.Chat.ID)
			}

			for _, hook := range b.hooks {
				if hook.Name == arg {
					t := fmt.Sprintf("%s: %s.", hook.Name, hook.Desc)
					b.sendText(t, input.Msg.Chat.ID)
					break
				}
			}
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

	b.sendText(t, input.Msg.Chat.ID)
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

		go b.handleCommand(a)
	case *HookInput:
		go b.handleHook(a)
	}
}

func (b *Bot) Run() {
	b.startTime = time.Now()

	log.Printf("Authorized on account %s", b.tg.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.tg.GetUpdatesChan(u)
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
