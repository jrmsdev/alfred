// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	fpath "path/filepath"

	_ "github.com/jrmsdev/alfred/internal/server/core/database"
	_ "github.com/jrmsdev/alfred/internal/server/core/router"

	"github.com/jrmsdev/alfred/internal/server"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const drvargs = "cache=shared&mode=rwc&_loc=UTC&_locking=EXCLUSIVE"

func main() {
	orm.RegisterDataBase("default", "sqlite3",
		fmt.Sprintf("file:%s?%s", fpath.FromSlash("./core.db"), drvargs))

	beego.BConfig.Listen.AdminPort = 27720
	beego.BConfig.WebConfig.AutoRender = false

	server.Start("core", "127.0.0.1:27719")
}
