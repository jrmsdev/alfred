// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package router

import (
	"github.com/jrmsdev/alfred/internal/server/api/controller"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v0",
		beego.NSRouter("/object", &controller.ObjectController{}, "GET:Get"),
		beego.NSRouter("/user", &controller.UserController{}, "GET:Get"),
	)
	beego.AddNamespace(ns)
	beego.Router("/status", &controller.Status{}, "GET:Get")
}
