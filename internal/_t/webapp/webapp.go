// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"

	"github.com/jrmsdev/alfred/internal/_t/check"
)

type testClient struct {
	Req *http.Request
	Resp *httptest.ResponseRecorder
}

func Client() *testClient {
	return &testClient{}
}

func (c *testClient) Get(url string) {
	c.Req, _ = http.NewRequest("GET", url, nil)
	c.Resp = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(c.Resp, c.Req)
}

func (c *testClient) Code(t *testing.T, expect int) {
	t.Helper()
	if check.NotEqual(t, c.Resp.Code, expect, "response code") {
		t.FailNow()
	}
}
