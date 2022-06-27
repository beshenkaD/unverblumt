package unverblumt

import (
	"os"
	"sync"

	"github.com/beshenkaD/unverblumt/internal/log"
	tb "gopkg.in/telebot.v3"
)

type Arg struct {
	Name     string
	Desc     string
	Required bool
}

type Command struct {
	Cmd     string
	Desc    string
	Args    []Arg
	Handler tb.HandlerFunc
}

type unverblumt struct {
	Bot *tb.Bot

	// sql.DB for disabling etc
	mut      sync.Mutex
	commands map[string]*Command
}

var instance *unverblumt

func Get() *unverblumt {
	return instance
}

func Run(m ...Module) {
	instance = new(unverblumt)

	s := tb.Settings{
		Token:     os.Getenv("UNVERBLUMT_TOKEN"),
		ParseMode: "HTML",
		OnError: func(e error, c tb.Context) {
			log.Warn.Println(e)
		},
	}

	var err error
	Get().Bot, err = tb.NewBot(s)
	if err != nil {
		log.Error.Fatalln(err)
	}
	Get().commands = make(map[string]*Command)

	loadBuiltins()
	loadModules(m)
	Get().setCommands()

	log.Info.Println("starting as: " + Get().Bot.Me.Username)
	Get().Bot.Start()
}

func (u *unverblumt) AddCommand(c *Command) {
	u.mut.Lock()
	defer u.mut.Unlock()

	u.commands[c.Cmd[1:]] = c
	u.Bot.Handle(c.Cmd, c.Handler)
}

func (u *unverblumt) setCommands() {

}

// add builtin middleware that rejects all disabled commands
func (u *unverblumt) disableLocal(cmd string, chatID int) {
}

func (u *unverblumt) disableGlobal(cmd string) {
}
