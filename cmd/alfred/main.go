// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/log"
)

func main() {
	rc := 0
	flags.Parse("alfred")
	log.Debug("init")

	action := flags.Options.Arg(0)
	if action == "" || action == "start" {
		rc = start()
	} else {
		log.Errorf("invalid action %s", action)
		rc = 1
	}

	log.Debug("exit(%d)", rc)
	os.Exit(rc)
}

func check(err error) {
	if err != nil {
		log.Error(err)
		os.Exit(2)
	}
}
