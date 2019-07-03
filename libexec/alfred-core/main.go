// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	_ "github.com/jrmsdev/alfred/internal/server/api/routers"

	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
	"github.com/jrmsdev/alfred"
)

func main() {
	rc := 0
	flags.Parse("alfred-core")
	log.Debug("init %s", alfred.Config.Core.Addr)
	server.Start(alfred.Config.Core.Addr)
	log.Debug("exit(%d)", rc)
	os.Exit(rc)
}
