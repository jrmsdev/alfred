// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"context"
	"sync"
	"time"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/core"
	"github.com/jrmsdev/alfred/internal/worker"
	"github.com/jrmsdev/alfred/log"
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

	if alfred.Config.Web.Addr != "" {
		log.Debug("web worker enabled")
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := worker.Start(ctx, "web")
			if err != nil {
				cancel()
				core.Stop()
			}
		}()
		log.Printf("http://%s/", alfred.Config.Web.Addr)
	} else {
		log.Debug("web worker not enabled")
	}

	wg.Wait()
	return 0
}
