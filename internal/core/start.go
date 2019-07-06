// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"fmt"
	fpath "path/filepath"

	_ "github.com/jrmsdev/alfred/internal/server/core/database"
	_ "github.com/jrmsdev/alfred/internal/server/core/router"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/server"
	"github.com/jrmsdev/alfred/log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const drvargs = "cache=shared&mode=rwc&_loc=UTC&_locking=EXCLUSIVE"

func Start() {
	log.Debug(alfred.Config.Core.Addr)

	orm.RegisterDataBase("default", "sqlite3",
		fmt.Sprintf("file:%s?%s",
			fpath.Join(alfred.Config.DataDir, "core.db"), drvargs))

	beego.BConfig.WebConfig.AutoRender = false

	server.Start("core", alfred.Config.Core.Addr)
}
