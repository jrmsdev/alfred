// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/astaxie/beego"
	"github.com/jrmsdev/alfred/log"

	_ "github.com/jrmsdev/alfred/internal/server/api/routers"
)

func start() {
	addr := "127.0.0.1:8180"
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.Run(addr)
}

func main() {
	log.Init("debug")
	log.Debug("start")
	start()
	log.Debug("end")
}
