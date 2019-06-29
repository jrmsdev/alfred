// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"

	"github.com/jrmsdev/alfred/log"
)

type Flags struct {
	Log struct {
		Level string
	}
}

var Config = new(Flags)

func init() {
	flag.StringVar(&Config.Log.Level, "log", "warn",
		"set log `level`: debug, warn, error or quiet")
}

func Parse() {
	flag.Parse()
	log.Init(Config.Log.Level)
}
