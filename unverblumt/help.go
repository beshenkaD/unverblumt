package unverblumt

import (
	"fmt"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"github.com/beshenkaD/unverblumt/internal/style"
)

func (u *unverblumt) GenerateHelp(lang string, args []string) string {
	if len(args) == 0 {
		return u._brief(lang)
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
	arg := func(lang string, a Arg) string {
		return i18n.TD(lang, "help", "%s : %s", a.Name, a.Desc)
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

		descriptions += arg(lang, a) + "\n    "
	}

	if required != "" {
		required = enquote(required, "{ ", "}")
		required = style.Underline(required)
	}

	if optional != "" {
		optional = enquote(optional, "[ ", "]")
		optional = style.Bold(optional)
	}

	return fmt.Sprintf("%s %s %s : %s\n    %s", c.Cmd, required, optional, c.Desc, descriptions)
}
