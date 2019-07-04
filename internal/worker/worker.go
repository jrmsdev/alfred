// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"context"
	"os/exec"
	fpath "path/filepath"
	"sync"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/web"
	"github.com/jrmsdev/alfred/log"
)

var Group = new(sync.WaitGroup)

func Start(ctx context.Context, name string) error {
	binfn := fpath.Join(alfred.Config.LibDir, "bin", name)
	if name == "web" {
		Group.Add(1)
		return web.Start(Group)
	}
	return dispatch(ctx, binfn)
}

func dispatch(ctx context.Context, binfn string) error {
	err := make(chan error)
	Group.Add(1)
	go worker(ctx, err, binfn)
	Group.Done()
	return <-err
}

func worker(ctx context.Context, err chan error, binfn string) {
	log.Debug("%s", binfn)
	cmd := exec.CommandContext(ctx, binfn)
	cmd.Start()
	err <- cmd.Wait()
}
