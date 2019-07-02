// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/alfred/internal/gen"
)

var data = map[string]string{
	"VMajor": "0",
	"VMinor": "0",
	"VPatch": "0",
}

func main() {
	gen.Template("version_info.go.in", &data)
}
