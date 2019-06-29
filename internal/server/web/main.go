package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jrmsdev/alfred/internal/server/web/routers"
)

func main() {
	beego.Run()
}
