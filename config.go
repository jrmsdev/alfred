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
	RunDir string
}

var Config *alfredConfig

func init() {
	Config = new(alfredConfig)
	Config.Log.Dir = getenv("ALFRED_LOGDIR",
		user.Home(".local", "alfred", "log"))
	Config.RunDir = getenv("ALFRED_RUNDIR",
		user.Home(".local", "alfred", "run"))
}

func getenv(varname, defval string) string {
	v := os.Getenv(varname)
	if v == "" {
		v = defval
	}
	return v
}
