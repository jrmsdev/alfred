// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"
)

func main() {
	log.Init("debug")
	log.Debug("start")
	log.Print("server start")
	server.Start()
	log.Debug("end")
}
