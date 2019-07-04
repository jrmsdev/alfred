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
	Group.Add(1)
	if name == "web" {
		return web.Start(Group)
	}
	return worker(ctx, binfn)
}

func worker(ctx context.Context, binfn string) error {
	log.Debug("%s", binfn)
	cmd := exec.CommandContext(ctx, binfn)
	cmd.Start()
	defer cmd.Process.Release()
	defer Group.Done()
	return nil
}
