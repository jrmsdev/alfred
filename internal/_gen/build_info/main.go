// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	fpath "path/filepath"
	"time"

	"github.com/jrmsdev/alfred/internal/_gen"
)

var data = map[string]string{
	"VersionBuild":  time.Now().UTC().Format("20060102.150405"),
	"InstallBinDir": getenv("ALFRED_BINDIR", fpath.FromSlash("/usr/local/bin")),
	"InstallLibDir": getenv("ALFRED_LIBDIR", fpath.FromSlash("/usr/local/lib/alfred")),
}

func getenv(varname, defval string) string {
	v := os.Getenv(varname)
	if v == "" {
		v = defval
	}
	return v
}

func main() {
	gen.Template("build_info.go.in", &data)
}
