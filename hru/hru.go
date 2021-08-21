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
	description := func() string {
		descs := []string{"Одебилевший", "Понадусеровый", "Щеневмерлый", "Майданутый", "Горелодупый", "Дебилоскачущий", "Бандерофашисткий"}

		t := descs[rand.Intn(len(descs))]
		if rand.Intn(2) == 1 {
			t1 := descs[rand.Intn(len(descs))]
			if t1 != t {
				t = t + " " + t1
			}
		}

		return t
	}

	name := func() string {
		names := []string{"швайнокарась", "свинопёс", "каклошвайн", "говнохуй", "бандера", "швайнопескарь", "хуйлопан"}
		names = append(names, []string{"cвин", "боров", "хряк", "борька", "хрюшка", "подсвинок", "свинтус", "поросенок"}...)

		return names[rand.Intn(len(names))]
	}

	action := func() string {
		actions := []string{"хрюкнул", "пукнул", "громко прохрюкал", "чавкает", "подхрюкивает", "подпукивает"}
		actions = append(actions, []string{"издает звуки хрюканья", "протестует", "срет"}...)

		for _, a := range actions {
			actions = append(actions, "снова"+" "+a)
			actions = append(actions, "опять"+" "+a)
		}

		return actions[rand.Intn(len(actions))]
	}

	emoji := func() string {
		emojis := []string{"🐖", "🐷", ")"}

		t := ""
		for i := 0; i < rand.Intn(10); i++ {
			t += emojis[rand.Intn(len(emojis))]
		}

		return t
	}

	return description() + " " + name() + " " + action() + " " + emoji()
}

func HruCommand(in *bot.CommandInput) (*bot.Output, error) {
	in.Bot.DeleteMessage(in.Msg.Chat.ID, in.Msg.MessageID)

	return &bot.Output{
		Text:    getHru(),
		UseHTML: false,
	}, nil
}

func HruWatcher(in *bot.HookInput) (*bot.Output, error) {
	validHru := regexp.MustCompile(`.*(х+р+ю+|h+r+u+)+.*`)

	if validHru.MatchString(strings.ToLower(in.Msg.Text)) {
		return &bot.Output{
			Text:    getHru(),
			UseHTML: false,
		}, nil
	}

	return nil, nil
}
