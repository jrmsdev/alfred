// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package web

import (
	"sync"

	"github.com/jrmsdev/alfred/log"
)

func Start(wg *sync.WaitGroup) error {
	log.Debug("web worker")
	defer wg.Done()
	return nil
}
