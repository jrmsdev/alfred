package main

import (
	_ "github.com/jrmsdev/alfred/internal/server/web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

