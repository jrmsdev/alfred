// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"context"
	//~ "os/exec"
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
	return dispatch(binfn)
}

func dispatch(binfn string) error {
	err := make(chan error)
	log.Debug("dispatch %s", binfn)
	Group.Add(1)
	go func() {
		err <- nil
	}()
	Group.Done()
	return <-err
}
