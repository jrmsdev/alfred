// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/log"
)

type Flags struct {
	Log struct {
		Level string
	}
}

var showVersion = false
var progname = filepath.Base(os.Args[0])

var Config = new(Flags)
var Options = flag.NewFlagSet(progname, flag.ExitOnError)

func init() {
	Options.BoolVar(&showVersion, "version", false, "show version and exit")
	Options.StringVar(&Config.Log.Level, "log", "warn",
		"set log `level`: debug, warn, error or quiet")
}

func parseArgs(args []string) {
	Options.Parse(args)
	log.Init(Config.Log.Level)
}

func Parse() {
	parseArgs(os.Args[1:])
	if showVersion {
		fmt.Println(alfred.Version(progname))
		os.Exit(0)
	}
}
