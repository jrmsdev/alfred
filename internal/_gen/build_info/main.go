// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred/internal/_gen"
)

var data = map[string]string{}

func main() {
	data["InstallBinDir"] = getenv("ALFRED_BINDIR",
		fpath.FromSlash("/usr/local/bin"))
	data["InstallLibDir"] = getenv("ALFRED_LIBDIR",
		fpath.FromSlash("/usr/local/lib/alfred"))
	gen.Template("build_info.go.in", &data)
}

func getenv(varname, defval string) string {
	v := os.Getenv(varname)
	if v == "" {
		v = defval
	}
	return v
}
