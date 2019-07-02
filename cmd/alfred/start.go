// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"sync"

	"github.com/jrmsdev/alfred/log"
)

func start() int {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	core(wg)
	web(wg)
	return 0
}

func core(wg *sync.WaitGroup) {
	log.Debug("core worker")
	defer wg.Done()
}

func web(wg *sync.WaitGroup) {
	log.Debug("web worker")
	defer wg.Done()
}
