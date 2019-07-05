// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package server

import (
	"fmt"
	fpath "path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/jrmsdev/alfred"
)

func Start(name, addr string) {
	logs.Async()
	logs.SetLogger(logs.AdapterFile,
		fmt.Sprintf(`{"filename":"%s.log","level":%d,"perm":"0640"}`,
			fpath.Join(alfred.Config.Log.Dir, name), logs.LevelDebug))
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.Run(addr)
}
