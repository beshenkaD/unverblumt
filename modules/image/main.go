package image

import (
	"math/rand"
	"strings"
	"time"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"github.com/beshenkaD/unverblumt/internal/log"
	u "github.com/beshenkaD/unverblumt/unverblumt"
	"gopkg.in/telebot.v3"
)

type Image struct{}

func (m *Image) Init() error {
	rand.Seed(time.Now().UnixNano())

	log.Info.Println("initialized module `image`")
	return nil
}

func (m *Image) Commands() []*u.Command {
	return []*u.Command{
		{
			Cmd:  "/image",
			Desc: "search for an image on the internet",
			Args: []u.Arg{
				{
					Name:     "keywords",
					Desc:     "search keywords",
					Required: true,
				},
			},
			Handler: image,
		},
		{
			Cmd:  "/image_unsafe",
			Desc: "like /image but without safe search option",
			Args: []u.Arg{
				{
					Name:     "keywords",
					Desc:     "search keywords",
					Required: true,
				},
			},
			Handler: imageUnsafe,
		},
	}
}

func _image(c telebot.Context, safe bool) error {
	if len(c.Args()) == 0 {
		return c.Reply(i18n.TC(c, "please provide keywords for search"))
	}

	imgs, err := searxImages(strings.Join(c.Args(), " "), safe)
	if err != nil {
		return err
	}

	if len(imgs) == 0 {
		return c.Reply(i18n.TC(c, "no images found for the given keywords"))
	}

	for {
		file := telebot.Photo{File: telebot.FromURL(imgs[rand.Intn(len(imgs))])}
		_, err := file.Send(c.Bot(), c.Recipient(), &telebot.SendOptions{
			ReplyTo: c.Message(),
		})

		if err != nil {
			continue
		}

		return nil
	}
}

func image(c telebot.Context) error {
	return _image(c, true)
}

func imageUnsafe(c telebot.Context) error {
	return _image(c, false)
}
