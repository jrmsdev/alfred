// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	_ "github.com/jrmsdev/alfred/internal/server/web/routers"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func main() {
	rc := 0
	flags.Parse("alfred-web")
	log.Debug("init %s", alfred.Config.Web.Addr)
	server.Start("web", alfred.Config.Web.Addr)
	log.Debug("exit(%d)", rc)
	os.Exit(rc)
}
