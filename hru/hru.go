package hru

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/beshenkaD/unverblumt/bot"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getHru() string {
	hruMap := make(map[int]string)

	list := []string{
		"Свинья хрюкнула)",
		"Свинтус похрюкивает)",
		"Понадусеровый швайнокарась опять расхрюкался))",
		"Горелодупый свинопёс харош хрюкать)",
		"Ну ты и хрюкнуло свиноеблое чудище)",
		"Не перди, свинка)",
		"Хряк расхрюкался)",
		"Боров похрюкивает)",
		"Угаманись, свинка)",
	}

	for i, v := range list {
		hruMap[1<<1+i] = v
	}

	return hruMap[rand.Intn(len(list))]
}

func HruCommand(in *bot.CommandInput) (*bot.Output, error) {
	in.Bot.DeleteMessage(in.Msg.Chat.ID, in.Msg.MessageID)

	return &bot.Output{
		Text:    getHru(),
		UseHTML: false,
	}, nil
}

func HruWatcher(in *bot.HookInput) (*bot.Output, error) {
	validHru := regexp.MustCompile(`(х+р+ю+)|(h+r+u+)`)

	if validHru.MatchString(strings.ToLower(in.Msg.Text)) {
		return &bot.Output{
			Text:    getHru(),
			UseHTML: false,
		}, nil
	}

	return nil, nil
}
