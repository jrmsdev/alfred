// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"os"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/log"
)

var showVersion = false
var Options *flag.FlagSet

func newOptions(progname string) {
	Options = flag.NewFlagSet(progname, flag.ExitOnError)

	Options.BoolVar(&showVersion, "version", false, "show version and exit")

	Options.StringVar(&alfred.Config.Log.Level, "log",
		alfred.Config.Log.Level,
		"set log `level`: default, debug, warn, error or quiet")

	Options.StringVar(&alfred.Config.Log.Dir, "logdir",
		alfred.Config.Log.Dir, "set log directory `path`")

	Options.StringVar(&alfred.Config.CfgDir, "cfgdir",
		alfred.Config.CfgDir, "set config directory `path`")

	Options.StringVar(&alfred.Config.RunDir, "rundir",
		alfred.Config.RunDir, "set runtime directory `path`")

	Options.StringVar(&alfred.Config.DataDir, "datadir",
		alfred.Config.DataDir, "set data directory `path`")

	Options.StringVar(&alfred.Config.CacheDir, "cachedir",
		alfred.Config.CacheDir, "set cache directory `path`")

	Options.StringVar(&alfred.Config.Core.Addr, "core.addr",
		alfred.Config.Core.Addr, "set core bind `address`")

	Options.StringVar(&alfred.Config.Web.Addr, "web.addr",
		alfred.Config.Web.Addr, "set web bind `address`")
}

func Parse(progname string) {
	newOptions(progname)
	Options.Parse(os.Args[1:])
	log.Init(alfred.Config.Log.Level)
	if showVersion {
		log.Printf("%s %s (%s)", progname, alfred.Version(), alfred.VersionBuild())
		os.Exit(0)
	}
	log.Printf("%s %s", progname, alfred.Version())
	alfred.Config.Validate()
}
