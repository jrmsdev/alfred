// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package routers

import (
	"github.com/jrmsdev/alfred/internal/server/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v0",
		beego.NSRouter("/object",
			&controllers.ObjectController{},
		),
		beego.NSRouter("/user",
			&controllers.UserController{},
		),
	)
	beego.AddNamespace(ns)
}
