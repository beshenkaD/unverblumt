// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package args

import (
	"errors"
	"flag"
	"strings"

	"github.com/beshenkaD/unverblumt/internal/log"
	"github.com/itzg/go-flagsfiller"
)

/*
   Splits string respecting quotes
*/
func split(s, sep string) (a []string) {
	sb := &strings.Builder{}
	quoted := false

	for _, r := range s {
		if r == '"' {
			quoted = !quoted
		} else if !quoted && r == ' ' {
			a = append(a, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(r)
		}
	}

	if sb.Len() > 0 {
		a = append(a, sb.String())
	}

	return a
}

/*
   Parses given string like: "/command -arg argument"
   see: https://pkg.go.dev/github.com/itzg/go-flagsfiller for more info
*/
func Parse(text string, dest interface{}) error {
	filler := flagsfiller.New()

	set := flag.FlagSet{}
	err := filler.Fill(&set, dest)

	if err != nil {
		log.Error.Println(err)
		return nil
	}

	if strings.Count(text, `"`)%2 != 0 {
		return errors.New("quotes count is not even")
	}

	return set.Parse(split(text, " ")[1:])
}
