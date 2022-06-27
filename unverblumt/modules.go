package unverblumt

import "github.com/beshenkaD/unverblumt/internal/log"

type Module interface {
	Init() error
	Commands() []*Command
}

func loadModules(modules []Module) {
	for _, m := range modules {
		err := m.Init()
		if err != nil {
			log.Warn.Println(err)
			break
		}

		for _, c := range m.Commands() {
			Get().AddCommand(c)
		}
	}
}
