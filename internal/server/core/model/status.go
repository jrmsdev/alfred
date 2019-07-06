// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Status struct {
	Id int
	Status string `orm:"size(50)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func GetStatus() string {
	db := orm.NewOrm()
	cur := &Status{Id: 0}
	err := db.Read(cur)
	if err != nil {
		return "error"
	}
	return "init"
}
