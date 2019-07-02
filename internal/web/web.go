// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package web

import (
	"github.com/jrmsdev/alfred/internal/worker"
	"github.com/jrmsdev/alfred/log"
)

func Start(wg *worker.Group) {
	log.Debug("web worker")
	defer wg.Done()
}
