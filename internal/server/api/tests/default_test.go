// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"net/http"
	"net/http/httptest"
	fpath "path/filepath"
	"runtime"
	"testing"

	"github.com/jrmsdev/alfred/internal/_t/check"

	_ "github.com/jrmsdev/alfred/internal/server/api/routers"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := fpath.Abs(fpath.Dir(fpath.Join(file, ".."+string(fpath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v0/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//~ beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	if check.NotEqual(t, w.Code, 200, "response status") {
		t.FailNow()
	}
}
