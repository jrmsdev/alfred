// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	fpath "path/filepath"
	"testing"

	"github.com/jrmsdev/alfred/internal/_t/webapp"

	_ "github.com/jrmsdev/alfred/internal/server/web/routers"

	"github.com/astaxie/beego"
)

func init() {
	apppath, _ := fpath.Abs(fpath.FromSlash("./.."))
	beego.TestBeegoInit(apppath)
}

func TestGet(t *testing.T) {
	c := webapp.Client()
	c.Get("/")
	c.Code(t, 200)
}
