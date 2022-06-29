package text

import (
	"strings"
	"sync"

	"github.com/beshenkaD/unverblumt/internal/i18n"
	"gopkg.in/telebot.v3"
)

func flip(c telebot.Context) error {
	if len(c.Args()) == 0 {
		return c.Reply(i18n.T(c.Sender().LanguageCode, "please provide string to flip"))
	}

	mutex.RLock()
	defer mutex.RUnlock()

	var r = ""
	for _, c := range strings.ToLower(strings.Join(c.Args(), " ")) {
		if n, ok := table[string(c)]; ok {
			r += n
			continue
		}
		r += string(c)
	}

	return c.Reply(reverse(r))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var mutex sync.RWMutex
var table = map[string]string{
	// en letters
	"z": "z",
	"y": "ʎ",
	"x": "x",
	"w": "ʍ",
	"v": "ʌ",
	"u": "n",
	"t": "ʇ",
	"s": "s",
	"r": "ɹ",
	"q": "b",
	"p": "d",
	"o": "o",
	"n": "u",
	"m": "ɯ",
	"k": "ʞ",
	"j": "ɾ",
	"i": "ᴉ",
	"h": "ɥ",
	"g": "ƃ",
	"f": "ɟ",
	"e": "ǝ",
	"d": "p",
	"c": "ɔ",
	"b": "q",
	"a": "ɐ",
	// ru letters
	"я": "ʁ",
	"ю": "oı",
	"э": "є",
	"ь": "q",
	"ы": "ıq",
	"ъ": "q",
	"щ": "m",
	"ш": "m",
	"ц": "ǹ",
	"х": "х",
	"ф": "ф",
	"у": "ʎ",
	"т": "ɯ",
	"с": "ɔ",
	"р": "d",
	"п": "u",
	"о": "о",
	"н": "н",
	"м": "w",
	"л": "v",
	"к": "ʞ",
	"й": "ņ",
	"и": "и",
	"з": "ε",
	"ж": "ж",
	"ё": "ǝ",
	"е": "ǝ",
	"д": "ɓ",
	"г": "ɹ",
	"в": "ʚ",
	"б": "ƍ",
	"а": "ɐ",
	// numbers
	"0": "0",
	"9": "6",
	"8": "8",
	"7": "ㄥ",
	"6": "9",
	"5": "ϛ",
	"4": "ㄣ",
	"3": "Ɛ",
	"2": "ᄅ",
	"1": "Ɩ",
	// punctuation
	"_":  "‾",
	")":  "(",
	"(":  ")",
	"&":  "⅋",
	",":  "'",
	".":  "˙",
	"'":  ",",
	"\"": ",,",
	"?":  "¿",
	"!":  "¡",
}
