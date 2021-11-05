// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package stat

import (
	"fmt"
	"strings"
	"sync"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"
)

var (
	start      = time.Now()                // Start time. Used for uptime
	processed  = 0                         // number of processed messages
	activeLock = sync.Mutex{}              // map is not thread-safe
	active     = make(map[int64]time.Time) // latest message timestamp by chat ID
	history    []string                    // last 1500 commands
)

// TODO: terminate this safely
func init() {
	go func() {
		for {
			activeLock.Lock()

			for chat, last := range active {
				if time.Since(last) >= 1*time.Minute {
					delete(active, chat)
				}
			}

			activeLock.Unlock()

			time.Sleep(1 * time.Minute)
		}
	}()
}

func collector(u *tb.Update) bool {
	if u.Message == nil {
		return true
	}

	activeLock.Lock()
	defer activeLock.Unlock()
	active[u.Message.Chat.ID] = time.Now()

	processed++

	if u.Message.Text[0] == '/' && len(u.Message.Text) > 1 {
		command := strings.Split(u.Message.Text, " ")[0]

		if len(history) == 1500 {
			history = []string{command}
		} else {
			history = append(history, command)
		}
	}

	return true
}

func stat(c tb.Context) error {
	const f = `<u>Uptime:</u> <b>%s</b>
<u>Messages processed:</u> <b>%d</b>
<u>Active chats:</u> <b>%d</b>
<u>Most used commands:</u> <b>%s</b>`

	activeLock.Lock()
	defer activeLock.Unlock()

	mostUsed := []string{}
	for _, i := range topMessages(5) {
		if i.Count > 0 {
			mostUsed = append(mostUsed, fmt.Sprintf("%s (%d)", i.Message, i.Count))
		}
	}

	return c.Send(fmt.Sprintf(
		f,
		time.Since(start).Truncate(time.Second).String(),
		processed,
		len(active),
		strings.Join(mostUsed, ", "),
	))
}

func getDuplicates(s []string) map[string]int {
	duplicate_frequency := make(map[string]int)

	for _, item := range s {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1
		} else {
			duplicate_frequency[item] = 1
		}
	}

	return duplicate_frequency
}

type item struct {
	Message string
	Count   int
}

func topMessages(count int) (s []item) {
	d := getDuplicates(history)

	for i := 0; i < count; i++ {
		current := ""
		prev := 0

		for str, dups := range d {
			if dups > prev {
				prev = dups
				current = str
			}
		}

		s = append(s, item{current, prev})
		delete(d, current)
	}

	return
}
