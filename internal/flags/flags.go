// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/log"
)

var showVersion = false
var Options *flag.FlagSet

func Parse(progname string) {
	newOptions(progname)
	parseArgs(os.Args[1:])
	if showVersion {
		fmt.Println(alfred.Version(progname))
		os.Exit(0)
	}
}

func newOptions(progname string) {
	Options = flag.NewFlagSet(progname, flag.ExitOnError)
	Options.BoolVar(&showVersion, "version", false, "show version and exit")

	Options.StringVar(&alfred.Config.Log.Level, "log", "warn",
		"set log `level`: debug, warn, error or quiet")
	Options.StringVar(&alfred.Config.Log.Dir, "logdir", alfred.Config.Log.Dir,
		"set log `directory` path")

	Options.StringVar(&alfred.Config.RunDir, "rundir", alfred.Config.RunDir,
		"set runtime `directory` path")
}

func parseArgs(args []string) {
	Options.Parse(args)
	log.Init(alfred.Config.Log.Level)
}
