// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	fpath "path/filepath"
	"sync"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/internal/web"
	"github.com/jrmsdev/alfred/log"
)

var Group = new(sync.WaitGroup)

func Start(name string) error {
	binfn := fpath.Join(alfred.Config.LibDir, "bin", name)
	Group.Add(1)
	if name == "web" {
		return web.Start(Group)
	}
	return dispatch(binfn)
}

func dispatch(binfn string) error {
	log.Debug("dispatch %s", binfn)
	Group.Done()
	return nil
}
