package quote

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beshenkaD/unverblumt/bot"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	fontSize = 20
	// Image must always be 700 pixels wide, but height may vary
	width     = 700
	minHeight = 400
)

var (
	face = getFontFace(fontSize)
)

func getFontFace(size int) font.Face {
	var font *truetype.Font

	font, _ = truetype.Parse(goregular.TTF)

	face := truetype.NewFace(font, &truetype.Options{Size: float64(size)})

	return face
}

func getPhotoPoint(height int) (x, y int) {
	return width / 6, height / 2
}

func getNamePoint(height int) (x, y int) {
	return 15, height - 15
}

func getStringWidth(s string) int {
	w := 0
	for _, r := range s {
		_, a, _ := face.GlyphBounds(r)
		w += a.Round()
	}

	return w
}

func getLinesHeight(lines []string) int {
	h := (len(lines) * (face.Metrics().Height.Ceil() + 2))
	if h < minHeight {
		return minHeight
	}

	return h
}

func getLines(s string, w int) []string {
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		var newLine string
		for _, word := range strings.Split(line, " ") {
			if getStringWidth(newLine+" "+word) > (width - w - 10) {
				lines = append(lines, newLine)
				newLine = word
			} else {
				newLine = newLine + " " + word
			}
		}

		newLine = strings.TrimSpace(newLine)
		if len(newLine) != 0 {
			lines = append(lines, newLine)
		}
		newLine = ""
	}

	lines[0] = "«" + lines[0]
	lines[len(lines)-1] = lines[len(lines)-1] + "»"

	return lines
}

func getName(firstName, lastName string, self bool) string {
	s := ""

	if self {
		s = "(Self Signed)"
	}

	return strings.Title(firstName) + " " + strings.Title(lastName) + " (c) " + s
}

type quoteMode int

const (
	lightMode = iota
	darkMode
)

func generateQuote(photo image.Image, firstName, lastName, quote string, self bool, mode quoteMode) interface{} {
	var (
		fg     color.Color
		bg     color.Color
		name   string
		drawer func()
	)

	const (
		textPointX = width / 3
	)

	lines := getLines(quote, textPointX)
	height := getLinesHeight(lines)

	dc := gg.NewContext(width, height)

	classic := func() {
		for i, line := range lines {
			y := height/2 - fontSize*len(lines)/2 + i*fontSize
			dc.DrawString(line, textPointX, float64(y))
		}
	}

	switch mode {
	case lightMode:
		fg = color.Black
		bg = color.White
		name = getName(firstName, lastName, self)
		drawer = classic
	case darkMode:
		fg = color.White
		bg = color.Black
		name = getName(firstName, lastName, self)
		drawer = classic
	default:
		fg = color.White
		bg = color.Black
		name = getName(firstName, lastName, self)
		drawer = classic
	}

	dc.SetFontFace(face)
	dc.SetColor(bg)
	dc.Clear()
	dc.SetColor(fg)

	// Draw quote text
	drawer()

	// Draw name
	nx, ny := getNamePoint(height)
	dc.DrawString(name, float64(nx), float64(ny))

	// Draw time
	t := time.Now().UTC().Format("02.01.2006 15:04")
	dc.DrawString(t, float64(nx+510), float64(ny))

	// Draw photo and make it round
	px, py := getPhotoPoint(height)
	dc.DrawEllipse(float64(px), float64(py), 100, 100)
	dc.Clip()
	dc.DrawImageAnchored(photo, px, py, 0.5, 0.5)

	i := dc.Image()

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, i, nil)

	IMG := tgbotapi.FileBytes{
		Name:  "quote123.jpeg",
		Bytes: buf.Bytes(),
	}

	return IMG
}

func Quote(in *bot.CommandInput) (*bot.Output, error) {
	var (
		firstName string
		lastName  string
		photo     image.Image
		text      string
		self      bool
	)

	if in.Msg.ReplyToMessage != nil {
		firstName = in.Msg.ReplyToMessage.From.FirstName
		lastName = in.Msg.ReplyToMessage.From.LastName
		text = in.Msg.ReplyToMessage.Text
		self = in.Msg.ReplyToMessage.From.ID == in.Msg.From.ID

		photos, _ := in.Bot.Tg.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{
			UserID: in.Msg.ReplyToMessage.From.ID,
			Limit:  1,
		})

		if len(photos.Photos) > 0 {
			f, err := in.Bot.Tg.GetFileDirectURL(photos.Photos[0][1].FileID)
			if err != nil {
				fmt.Println(err.Error())
			}

			resp, err := grab.Get(filepath.Join(os.TempDir(), f+".jpg"), f)
			if err != nil {
				log.Println(err.Error())
			}

			photo, _ = gg.LoadImage(resp.Filename)
			photo = resize.Resize(200, 200, photo, resize.Bilinear)
			os.Remove(resp.Filename)
		} else {
			ct := gg.NewContext(200, 200)

			photo = ct.Image()
		}
	} else {
		return nil, errors.New("error")
	}

	var (
		mode quoteMode
	)

	if len(in.Args) > 0 {
		switch in.Args[0] {
		case "dark":
			mode = darkMode
		case "light":
			mode = lightMode
		default:
			return nil, bot.ErrWrongArgs
		}
	}

	if len(text) == 0 {
		return nil, errors.New("empty quote")
	}

	quoteIMG := generateQuote(photo, firstName, lastName, text, self, mode)

	return &bot.Output{
		Text:  "",
		Photo: quoteIMG,
	}, nil
}
