// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package server

import (
	_ "github.com/jrmsdev/alfred/internal/server/api/routers"

	"github.com/astaxie/beego"
)

func Start() {
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.Run()
}
