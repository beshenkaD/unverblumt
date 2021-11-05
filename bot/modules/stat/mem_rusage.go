// Copyright (C) 2021 beshenkaD
// SPDX-License-Identifier: GPL-2.0-or-later

// +build linux darwin freebsd netbsd openbsd

package stat

import (
	"strconv"
	"syscall"

	"github.com/beshenkaD/unverblumt/internal/log"
)

func init() {
	getRusage = rusage
}

func rusage() string {
	var usage syscall.Rusage

	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		log.Error.Println(err)
		return ""
	}

	return strconv.Itoa(int(usage.Maxrss / 1024))
}
