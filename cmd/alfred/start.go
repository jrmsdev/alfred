// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"context"
	"os"

	"github.com/jrmsdev/alfred/internal/worker"
	"github.com/jrmsdev/alfred/log"
)

var bgctx = context.Background()

func start() int {
	ctx, cancel := context.WithCancel(bgctx)
	defer cancel()
	check(worker.Start(ctx, "core"), cancel)
	check(worker.Start(ctx, "web"), cancel)
	worker.Group.Wait()
	return 0
}

func check(err error, cancel context.CancelFunc) {
	if err != nil {
		cancel()
		log.Error(err)
		os.Exit(2)
	}
}
