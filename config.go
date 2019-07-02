// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	"os"

	"github.com/jrmsdev/alfred/internal/os/user"
)

type alfredConfig struct {
	Log struct {
		Level string
		Dir   string
	}
	Dir string
	RunDir string
	DataDir string
	CacheDir string
}

var Config *alfredConfig

func init() {
	Config = new(alfredConfig)
	Config.Log.Dir = getenv("ALFRED_LOGDIR",
		user.Home(".local", "alfred", "log"))
	Config.Dir = getenv("ALFRED_CFGDIR",
		user.Home(".config", "alfred"))
	Config.RunDir = getenv("ALFRED_RUNDIR",
		user.Home(".local", "alfred", "run"))
	Config.DataDir = getenv("ALFRED_DATADIR",
		user.Home(".local", "alfred", "data"))
	Config.CacheDir = getenv("ALFRED_CACHEDIR",
		user.Home(".cache", "alfred"))
}

func getenv(varname, defval string) string {
	v := os.Getenv(varname)
	if v == "" {
		v = defval
	}
	return v
}
