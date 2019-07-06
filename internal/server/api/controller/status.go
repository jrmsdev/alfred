// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package controller

import (
	"github.com/astaxie/beego"

	"github.com/jrmsdev/alfred/internal/server/api/model"
)

type Status struct {
	beego.Controller
}

func (s *Status) Get() {
	s.Data["json"] = beego.M{"status": model.GetStatus()}
	s.ServeJSON()
}
