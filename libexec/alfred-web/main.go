// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	_ "github.com/jrmsdev/alfred/internal/server/web/routers"

	"github.com/jrmsdev/alfred/internal/flags"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func main() {
	flags.Parse("alfred-web")
	log.Debug("start")
	server.Start("127.0.0.1:21680")
	log.Debug("end")
}
