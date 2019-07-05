package main

import (
	_ "github.com/jrmsdev/alfred/internal/server/web/routers"

	"github.com/jrmsdev/alfred/internal/server"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.AdminPort = 21681
	beego.BConfig.WebConfig.AutoRender = true
	server.Start("web", "127.0.0.1:21680")
}
