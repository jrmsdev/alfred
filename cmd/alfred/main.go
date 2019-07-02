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
	action := flags.Options.Arg(1)
	if action == "" || action == "start" {
		rc = start()
	}
	log.Debug("exit(%d)", rc)
	os.Exit(rc)
}
