// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package web

import (
	"sync"

	"github.com/jrmsdev/alfred/log"
)

func Start(wg *sync.WaitGroup) {
	log.Debug("web worker")
	defer wg.Done()
}
