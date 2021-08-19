package openweathermap

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/beshenkaD/unverblumt/bot"
	owm "github.com/briandowns/openweathermap"
)

var (
	token string
)

func init() {
	token = os.Getenv("BOT_OWM")
}

const weatherTemplate = `Current weather for <b>{{.Name}}</b>:
    Conditions: {{range .Weather}} <code>{{.Description}}</code> {{end}}
    Now:         <code>{{.Main.Temp}} °C</code>
    High:        <code>{{.Main.TempMax}} °C</code>
    Low:         <code>{{.Main.TempMin}} °C</code>
`

func Weather(in *bot.CommandInput) (*bot.Output, error) {
	var (
		where = "Moscow"
		lang  = "en"
	)

	if len(in.Args) > 0 {
		where = in.Args[0]
	}
	if len(in.Args) > 1 {
		lang = in.Args[1]
	}

	w, err := owm.NewCurrent("C", lang, token)
	if err != nil {
		return nil, err
	}

	err = w.CurrentByName(where)
	if err != nil {
		return nil, fmt.Errorf("cannot find such city: %s", where)
	}

	tmpl, _ := template.New("weather").Parse(weatherTemplate)
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, w)

	return &bot.Output{
		Text:    buf.String(),
		UseHTML: true,
		Photo:   nil,
	}, nil
}
