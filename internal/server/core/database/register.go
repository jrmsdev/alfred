// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package database

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jrmsdev/alfred/internal/server/core/model"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(&model.Status{})
}
