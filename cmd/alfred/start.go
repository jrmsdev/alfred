// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/alfred/internal/worker"
)

func start() int {
	check(worker.Start("core"))
	check(worker.Start("web"))
	worker.Group.Wait()
	return 0
}
