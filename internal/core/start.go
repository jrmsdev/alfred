// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	_ "github.com/jrmsdev/alfred/internal/server/core/database"
	_ "github.com/jrmsdev/alfred/internal/server/core/router"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"

	"github.com/astaxie/beego"
)

func Start() {
	log.Debug(alfred.Config.Core.Addr)
	beego.BConfig.WebConfig.AutoRender = false
	server.Start("core", alfred.Config.Core.Addr)
}
