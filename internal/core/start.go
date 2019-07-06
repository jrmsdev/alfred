// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	_ "github.com/jrmsdev/alfred/internal/server/core/router"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func Start() {
	log.Debug(alfred.Config.Core.Addr)
	server.Start("core", alfred.Config.Core.Addr)
}
