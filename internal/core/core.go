// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/jrmsdev/alfred/internal/worker"
	"github.com/jrmsdev/alfred/log"
)

func Start(wg *worker.Group) {
	log.Debug("core worker")
	defer wg.Done()
}
