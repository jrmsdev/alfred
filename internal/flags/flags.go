// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/jrmsdev/alfred/log"
)

type Flags struct {
	Log struct {
		Level string
	}
}

var Config = new(Flags)
var Options = flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

func init() {
	Options.StringVar(&Config.Log.Level, "log", "warn",
		"set log `level`: debug, warn, error or quiet")
}

func parseArgs(args []string) {
	Options.Parse(args)
	log.Init(Config.Log.Level)
}

func Parse() {
	parseArgs(os.Args[1:])
}
