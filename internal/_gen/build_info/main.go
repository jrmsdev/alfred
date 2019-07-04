// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred/internal/_gen"
)

var data = map[string]string{
	"InstallPrefix": "",
}

func main() {
	prefix := os.Getenv("ALFRED_INSTALL_PREFIX")
	if prefix == "" {
		prefix = fpath.FromSlash("/usr/local")
	}
	data["InstallPrefix"] = prefix
	gen.Template("build_info.go.in", &data)
}
