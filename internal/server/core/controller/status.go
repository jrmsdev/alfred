// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package controller

import (
	"github.com/astaxie/beego"

	"github.com/jrmsdev/alfred/internal/server/core/model"
)

type Status struct {
	beego.Controller
}

func (s *Status) Get() {
	s.Data["json"] = model.GetStatus()
	s.ServeJSON()
}
