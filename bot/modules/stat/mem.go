// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

package stat

import (
	"fmt"
	"runtime"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"
)

var getRusage func() string = nil

func memory(c tb.Context) error {
	const f = `<u>Alloc:</u> <b>%d MB</b>
<u>Sys:</u> <b>%d MB</b>
<u>Heap in use:</u> <b>%d MB</b>
<u>Rusage:</u> <b>%s</b>
<u>Heap objects:</u> <b>%d</b>
<u>Last GC:</u> <b>%s</b>
`
	bToMB := func(b uint64) uint64 {
		return b / 1024 / 1024
	}

	r := "not available"
	if getRusage != nil {
		r = getRusage() + " MB"
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return c.Send(fmt.Sprintf(
		f,
		bToMB(m.Alloc),
		bToMB(m.Sys),
		bToMB(m.HeapInuse),
		r,
		m.HeapObjects,
		time.Since(time.Unix(0, int64(m.LastGC))).Round(time.Second)))
}
