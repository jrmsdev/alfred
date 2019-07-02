// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/internal/worker"
	"github.com/jrmsdev/alfred/log"
)

var wg *worker.Group

func main() {
	rc := 0
	wg = new(worker.Group)
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
