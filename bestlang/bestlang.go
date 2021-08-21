package bestlang

import (
	"math/rand"
	"regexp"

	"github.com/beshenkaD/unverblumt/bot"
)

func init() {

}

func getBestLanguage() string {
	l := []string{"haskell.", "common lisp.", "elixir.", "erlang.", "C.", "GO."}

	return l[rand.Intn(len(l))]
}

func BestLangWatcher(in *bot.HookInput) (*bot.Output, error) {
	lang := regexp.MustCompile(`.*(лучший язык|best language).*`)

	if lang.MatchString(in.Msg.Text) {
		return &bot.Output{
			Text:     getBestLanguage(),
			UseHTML:  false,
			Photo:    nil,
			PhotoURL: "",
		}, nil
	}

	return nil, nil
}
