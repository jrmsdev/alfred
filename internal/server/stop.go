// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package server

import (
	"github.com/astaxie/beego"
)

func Stop() error {
	return beego.BeeApp.Server.Close()
}
