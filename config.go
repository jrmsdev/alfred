// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	"os"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred/errors"
	"github.com/jrmsdev/alfred/internal/os/user"
	"github.com/jrmsdev/alfred/log"
)

var installBinDir = fpath.FromSlash("/usr/local/bin")
var installLibDir = fpath.FromSlash("/usr/local/lib/alfred")

type alfredConfig struct {
	validateError bool

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
	//~ Config.Web.Addr = getenv("ALFRED_WEB", "127.0.0.1:21680")
	Config.Web.Addr = getenv("ALFRED_WEB", "")

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

func (obj *alfredConfig) checkLogLevel() {
	valid := map[string]bool{
		"default": true,
		"debug":   true,
		"warn":    true,
		"error":   true,
		"quiet":   true,
	}
	if !valid[obj.Log.Level] {
		log.Errorf("invalid log level: %s", obj.Log.Level)
		obj.validateError = true
	}
}

func (obj *alfredConfig) checkIsAbs(name, checkpath string) {
	if !fpath.IsAbs(checkpath) {
		err := errors.New(errors.NotAbsPath, name, checkpath)
		log.Error(err)
		obj.validateError = true
	}
}

func (obj *alfredConfig) Validate() {
	log.Debug("validate")
	obj.validateError = false
	obj.checkLogLevel()
	obj.checkIsAbs("logdir", obj.Log.Dir)
	obj.checkIsAbs("cfgdir", obj.CfgDir)
	obj.checkIsAbs("rundir", obj.RunDir)
	obj.checkIsAbs("datadir", obj.DataDir)
	obj.checkIsAbs("cachedir", obj.CacheDir)
	obj.checkIsAbs("runtime bin dir", obj.BinDir)
	obj.checkIsAbs("runtime lib dir", obj.LibDir)
	if obj.validateError {
		os.Exit(1)
	}
	obj.preChecks()
}

func (obj *alfredConfig) checkError(err error) {
	if err != nil {
		log.Error(err)
		os.Exit(2)
	}
}

func (obj *alfredConfig) dirExists(path string) {
	f, err := os.Stat(path)
	obj.checkError(err)
	if !f.IsDir() {
		log.Errorf("%s is not a directory", path)
		os.Exit(2)
	}
}

func (obj *alfredConfig) preChecks() {
	obj.checkError(os.MkdirAll(obj.Log.Dir, 0750))
	obj.checkError(os.MkdirAll(obj.RunDir, 0750))
	obj.checkError(os.MkdirAll(obj.DataDir, 0750))
	obj.checkError(os.MkdirAll(obj.CacheDir, 0750))
	obj.dirExists(obj.LibDir)
}
