// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"sync"

	"github.com/jrmsdev/alfred/internal/core"
)

func start() int {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		core.Start()
		wg.Done()
	}()
	wg.Wait()
	return 0
}
