// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package server

import (
	"github.com/astaxie/beego"
)

func Start(addr string) {
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.Run(addr)
}
