// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	"os"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred/internal/os/user"
)

var installBinDir = fpath.FromSlash("/usr/local/bin")
var installLibDir = fpath.FromSlash("/usr/local/lib/alfred")

type alfredConfig struct {
	Log struct {
		Level string
		Dir   string
	}
	Core struct {
		Addr string
	}
	Web struct {
		Addr string
	}
	CfgDir   string
	RunDir   string
	DataDir  string
	CacheDir string
	BinDir   string
	LibDir   string
}

var Config *alfredConfig

func init() {
	Config = new(alfredConfig)
	Config.Log.Level = getenv("ALFRED_LOG", "default")

	Config.Log.Dir = getenv("ALFRED_LOGDIR",
		user.Home(".local", "alfred", "log"))

	Config.CfgDir = getenv("ALFRED_CFGDIR",
		user.Home(".config", "alfred"))

	Config.RunDir = getenv("ALFRED_RUNDIR",
		user.Home(".local", "alfred", "run"))

	Config.DataDir = getenv("ALFRED_DATADIR",
		user.Home(".local", "alfred", "data"))

	Config.CacheDir = getenv("ALFRED_CACHEDIR",
		user.Home(".cache", "alfred"))

	Config.Core.Addr = getenv("ALFRED_CORE", "127.0.0.1:27719")
	Config.Web.Addr = getenv("ALFRED_WEB", "127.0.0.1:21680")

	Config.BinDir = installBinDir
	Config.LibDir = installLibDir
}

func getenv(varname, defval string) string {
	v := os.Getenv(varname)
	if v == "" {
		v = defval
	}
	return v
}
