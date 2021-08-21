package cat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beshenkaD/unverblumt/bot"
)

func getCat() (string, error) {
	resp, err := http.Get("https://thatcopy.pw/catapi/rest/")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	catOut := struct {
		URL string `json:"url"`
	}{}

	err = json.Unmarshal(body, &catOut)

	return catOut.URL, err
}

func Cat(in *bot.CommandInput) (*bot.Output, error) {
	cat, err := getCat()
	if err != nil {
		return nil, err
	}

	return &bot.Output{
		Text:     "<b><i>meow :3</i></b>",
		UseHTML:  true,
		Photo:    nil,
		PhotoURL: cat,
	}, nil
}
