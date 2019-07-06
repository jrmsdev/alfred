// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Status struct {
	Id      int       `json:"-"`
	Status  string    `orm:"size(50)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Error   string    `orm:"-"`
}

func GetStatus() *Status {
	db := orm.NewOrm()
	cur := &Status{Id: 0}
	err := db.Read(cur)
	if err != nil {
		cur.Error = err.Error()
		cur.Status = "error"
	}
	return cur
}
