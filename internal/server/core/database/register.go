// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package database

import (
	"fmt"
	fpath "path/filepath"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/server/core/model"
)

const drvargs = "cache=shared&mode=rwc&_loc=UTC&_locking=EXCLUSIVE"

func init() {
	orm.RegisterModel(&model.Status{})

	orm.RegisterDataBase("default", "sqlite3",
		fmt.Sprintf("file:%s?%s",
			fpath.Join(alfred.Config.DataDir, "core.db"), drvargs))
}
