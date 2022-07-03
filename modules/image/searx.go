package image

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Searx struct {
	Results []struct {
		Image string `json:"img_src"`
	} `json:"results"`
}

func searxImages(keywords string, safe bool) (r []string, e error) {
	const URL = "https://searx.zapashcanon.fr/search"

	reqURL := URL + "?q=" + url.QueryEscape(keywords) + "&format=json&categories=images"
	if safe {
		reqURL += "&safesearch=1"
	}

	req, _ := http.NewRequest("GET", reqURL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var s = Searx{}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	for _, i := range s.Results {
		if strings.HasPrefix(i.Image, "//") {
			i.Image = strings.Replace(i.Image, "//", "https://", 1)
		}

		r = append(r, i.Image)
	}

	return
}
