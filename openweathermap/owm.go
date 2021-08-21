package openweathermap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/beshenkaD/unverblumt/bot"
	owm "github.com/briandowns/openweathermap"
	"github.com/zsefvlol/timezonemapper"
)

var (
	token string
	icons map[string]string
)

func init() {
	token = os.Getenv("BOT_OWM")

	s := `
{ 
		"01d":"☀️" ,
		"02d":"⛅️" ,
		"03d":"☁️" ,
		"04d":"☁️" ,
		"09d":"\uD83C\uDF27" ,
		"10d":"\uD83C\uDF26" ,
		"11d":"⛈" ,
		"13d":"❄️" ,
		"50d":"\uD83C\uDF2B" ,
		"01n":"\uD83C\uDF11" ,
		"02n":"\uD83C\uDF11 ☁" ,
		"03n":"☁️" ,
		"04n":"️️☁☁" ,
		"09n":"\uD83C\uDF27" ,
		"10n":"☔️" ,
		"11n":"⛈" ,
		"13n":"❄️" ,
		"50n":"\uD83C\uDF2B"
}`

	json.Unmarshal([]byte(s), &icons)
}

func parse(args []string) (city, lang string) {
	city = "Moscow"
	lang = "en"

	if len(args) == 0 {
		return
	}

	if len(args) > 0 {
		city = args[0]
	}
	if len(args) > 1 {
		lang = args[1]
	}

	return
}

func getIcon(code string) string {
	return icons[code]
}

const weatherTemplate = `Current weather for <b>{{.Name}}</b>:<code>
    Conditions: {{range .Weather}}</code><b>{{.Description}}</b><code> {{getIcon .Icon}} {{end}}
    Temperature:
	    - Now:         {{.Main.Temp}} °C
	    - Feels like:  {{.Main.FeelsLike}} °C
    Humidity:    {{.Main.Humidity}}%
    Wind speed:  {{.Wind.Speed}} m/s
    Cloudiness:  {{.Clouds.All}}%
</code>
`
const forecastTemplate = `Weather Forecast for <b>{{.City.Name}}</b>:<code>
{{range .List}} </code><u>{{formatDate .Dt}}</u><code>:
    - Conditions:  {{range .Weather}}</code><b>{{.Description}}</b><code> {{getIcon .Icon}} {{end}}
    - Temp:        {{.Main.Temp}} °C
    - Feels like:  {{.Main.FeelsLike}} °C

{{end}}
</code>
`

func Weather(in *bot.CommandInput) (*bot.Output, error) {
	city, lang := parse(in.Args)

	w, err := owm.NewCurrent("C", lang, token)
	if err != nil {
		return nil, err
	}

	err = w.CurrentByName(city)
	if err != nil {
		return nil, fmt.Errorf("cannot find such city: %s", city)
	}

	tmpl, _ := template.New("weather").Funcs(template.FuncMap{
		"getIcon": getIcon,
	}).Parse(weatherTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, w)

	return &bot.Output{
		Text:    buf.String(),
		UseHTML: true,
		Photo:   nil,
	}, nil
}

func Forecast(in *bot.CommandInput) (*bot.Output, error) {
	city, lang := parse(in.Args)

	w, err := owm.NewForecast("5", "C", lang, token)
	if err != nil {
		return nil, err
	}

	err = w.DailyByName(city, 5)
	if err != nil {
		return nil, err
	}

	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	if forecast.City.Name == "" {
		return nil, fmt.Errorf("cannot find such city: %s", city)
	}

	tmpl, _ := template.New("forecast").Funcs(template.FuncMap{
		"getIcon": getIcon,
		"formatDate": func(dt int) string {
			var lat = forecast.City.Coord.Latitude
			var lon = forecast.City.Coord.Longitude

			loc, _ := time.LoadLocation(timezonemapper.LatLngToTimezoneString(lat, lon))

			t := time.Unix(int64(dt), 0).In(loc)
			return t.Format("02 Jan. Monday. 15:04 MST")
		},
	}).Parse(forecastTemplate)

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, forecast)

	return &bot.Output{
		Text:    buf.String(),
		UseHTML: true,
		Photo:   nil,
	}, nil
}
