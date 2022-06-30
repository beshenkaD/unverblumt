package random

import (
	"time"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"github.com/beshenkaD/unverblumt/internal/style"
	"gopkg.in/telebot.v3"
)

func crystalBall(c telebot.Context) error {
	var answers = []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes — definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Signs point to yes",
		"Yes",
		"Reply hazy, try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don’t count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	var code = c.Sender().LanguageCode

	var shaking = i18n.T(code, "shaking the magic ball")
	msg, _ := c.Bot().Reply(c.Message(), shaking)

	for i := 0; i < randIntInRange(3, 10); i++ {
		shaking += "."
		c.Bot().Edit(msg, shaking)
		time.Sleep(1 * time.Second)
	}

	var said = i18n.T(code, "magic ball said: ")
	_, err := c.Bot().Edit(msg, said+style.Enquote(i18n.T(code, answers[myRand.Intn(20)])))

	return err
}
