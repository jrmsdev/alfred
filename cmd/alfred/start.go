// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/alfred/internal/core"
	"github.com/jrmsdev/alfred/internal/web"
)

func start() int {
	wg.Add(2)
	core.Start(wg)
	web.Start(wg)
	return 0
}
