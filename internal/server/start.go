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
	beego.BConfig.AppName = fmt.Sprintf("alfred-%s", name)
	beego.BConfig.ServerName = beego.BConfig.AppName
	beego.BConfig.RecoverPanic = true
	beego.BConfig.EnableErrorsShow = true
	beego.BConfig.Log.EnableStaticLogs = true

	// web config
	beego.BConfig.WebConfig.EnableDocs = false
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.BConfig.WebConfig.FlashName = fmt.Sprintf("alfred_%s_flash", name)
	beego.BConfig.WebConfig.FlashSeparator = fmt.Sprintf("ALFRED%sFLASH", name)

	// http config
	beego.BConfig.Listen.EnableHTTP = true
	beego.BConfig.Listen.EnableHTTPS = false
	beego.BConfig.Listen.EnableAdmin = false
	beego.BConfig.Listen.AdminAddr = "127.0.0.1"

	// session config
	beego.BConfig.WebConfig.Session.SessionOn = false
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = fmt.Sprintf("alfred_%s_session", name)

	if beego.BConfig.RunMode == "dev" {
		// dev config
		logs.SetLogger(logs.AdapterConsole)
		beego.BConfig.Listen.EnableAdmin = true
	} else {
		// prod config
		logs.Async()
		logs.SetLogger(logs.AdapterFile,
			fmt.Sprintf(`{"filename":"%s.log","level":%d,"perm":"0640"}`,
				fpath.Join(alfred.Config.Log.Dir, name), logs.LevelDebug))
	}

	// run webapp
	beego.Run(addr)
}
