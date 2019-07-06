// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	_ "github.com/jrmsdev/alfred/internal/server/api/router"

	"github.com/jrmsdev/alfred/internal/server"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.AdminPort = 27720
	beego.BConfig.WebConfig.AutoRender = false
	server.Start("core", "127.0.0.1:27719")
}
