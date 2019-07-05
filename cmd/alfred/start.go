// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"context"
	"sync"
	"time"

	"github.com/jrmsdev/alfred/internal/core"
	"github.com/jrmsdev/alfred/internal/web"
)

var bgctx = context.Background()

func start() int {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		core.Start()
		wg.Done()
	}()
	time.Sleep(100 * time.Millisecond)
	ctx, cancel := context.WithCancel(bgctx)
	defer cancel()
	wg.Add(1)
	go func() {
		web.Start(ctx)
		wg.Done()
	}()
	wg.Wait()
	return 0
}
