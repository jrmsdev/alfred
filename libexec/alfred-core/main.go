// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	_ "github.com/jrmsdev/alfred/internal/server/api/routers"

	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func main() {
	rc := 0
	flags.Parse("alfred-core")
	log.Debug("init")
	server.Start("127.0.0.1:21600")
	log.Debug("exit(%d)", rc)
	os.Exit(rc)
}
