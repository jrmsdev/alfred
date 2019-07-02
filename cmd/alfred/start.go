// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"sync"

	"github.com/jrmsdev/alfred/log"
)

func start() int {
	wg := new(sync.WaitGroup)
	log.Debug("start")
	wg.Add(2)
	worker("alfred-core", wg)
	worker("alfred-web", wg)
	log.Debug("end")
	return 0
}

func worker(name string, wg *sync.WaitGroup) {
	log.Debug("%s worker", name)
	defer wg.Done()
}
