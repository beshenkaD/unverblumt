package unverblumt

import (
	"context"
	"time"

	"gopkg.in/telebot.v3"
)

type Animation struct {
	Frames       []string
	StartMessage string
	EndMessage   string
	Speed        time.Duration

	current int
}

func (a *Animation) Frame() string {
	if a.current == len(a.Frames) {
		a.current = 0
	}

	a.current++

	return a.Frames[a.current-1]
}

func GentooAnimationFrames() []string {
	return []string{
		"—", "\\", "|", "/",
	}
}

func DotsAnimationFrames() []string {
	return []string{
		".", "..", "...",
	}
}

func (u *unverblumt) ReplyWithAnimation(ctx context.Context, to *telebot.Message, a Animation) error {
	msg, _ := u.Bot.Reply(to, a.StartMessage)

	for {
		select {
		case <-ctx.Done():
			_, err := u.Bot.Edit(msg, a.EndMessage)
			return err
		default:
		}

		_, err := u.Bot.Edit(msg, a.StartMessage+a.Frame())
		if err != nil {
			return err
		}

		time.Sleep(a.Speed)
	}
}
