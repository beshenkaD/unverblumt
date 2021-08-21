package shortener

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/beshenkaD/unverblumt/bot"
	"mvdan.cc/xurls/v2"
)

var (
	apiURL   = "https://kutt.it/api/"
	urlRegex = xurls.Relaxed()
	token    = os.Getenv("BOT_KUTT")
)

type request struct {
	Target   string `json:"target"`
	ExpireIn string `json:"expire_in"`
	Password string `json:"password"`
}

type response struct {
	Link string `json:"link"`
}

func short(sr request) (string, error) {
	body, err := json.Marshal(sr)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL+"v2/links", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("X-API-Key", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return "", fmt.Errorf("kutt returned non-20* code: %d", resp.StatusCode)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}

	return r.Link, nil
}

func Short(in *bot.CommandInput) (*bot.Output, error) {
	if !urlRegex.MatchString(in.Msg.Text) {
		return nil, errors.New("you don't have URL's in your message")
	}

	//           0         1           2
	// /short <target> <expire_in> <password>

	if len(in.Args) < 1 {
		return nil, bot.ErrNoArgs
	}

	var sr request
	for i, arg := range in.Args {
		if i == 0 {
			sr.Target = arg
		}
		if i == 1 {
			sr.ExpireIn = arg
		}
		if i == 2 {
			sr.Password = arg
		}
	}

	link, err := short(sr)

	return &bot.Output{Text: link}, err
}

func ShortBunch(in *bot.CommandInput) (*bot.Output, error) {
	if len(in.Args) < 1 {
		return nil, bot.ErrNoArgs
	}

	// /shortBunch "example.com:expireIn:password" "example.com::password"... etc

	out := ""
	for i, arg := range in.Args {
		raw := strings.Split(arg, ":")
		sr := request{
			Target:   raw[0],
			ExpireIn: raw[1],
			Password: raw[2],
		}

		link, err := short(sr)
		if err != nil {
			out += fmt.Sprintf("%d. Failed: %s\n", i, err.Error())
		} else {
			out += fmt.Sprintf("%d. %s\n", i, link)
		}
	}

	return &bot.Output{
		Text: strings.TrimSpace(out),
	}, nil
}
