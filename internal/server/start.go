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
	// app config
	beego.BConfig.AppName = name
	beego.BConfig.ServerName = name
	beego.BConfig.RecoverPanic = true
	beego.BConfig.EnableErrorsShow = true
	// web config
	beego.BConfig.WebConfig.EnableDocs = false
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.BConfig.WebConfig.FlashName = fmt.Sprintf("alfred_%s_flash", name)
	// http config
	beego.BConfig.Listen.EnableHTTP = true
	beego.BConfig.Listen.EnableHTTPS = false
	beego.BConfig.Listen.EnableAdmin = false
	beego.BConfig.Listen.AdminAddr = "127.0.0.1"
	// session config
	beego.BConfig.WebConfig.Session.SessionOn = false
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = fmt.Sprintf("alfred_%s_session", name)
	// dev config
	if beego.BConfig.RunMode == "dev" {
		logs.SetLogger(logs.AdapterConsole)
		beego.BConfig.WebConfig.EnableDocs = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		//~ beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.Listen.EnableAdmin = true
	// prod config
	} else {
		logs.Async()
		logs.SetLogger(logs.AdapterFile,
			fmt.Sprintf(`{"filename":"%s.log","level":%d,"perm":"0640"}`,
				fpath.Join(alfred.Config.Log.Dir, name), logs.LevelDebug))
	}
	beego.Run(addr)
}
