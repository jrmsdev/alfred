// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func Stop() {
	log.Debug(alfred.Config.Core.Addr)
	err := server.Stop()
	if err != nil {
		log.Error(err)
	}
}
