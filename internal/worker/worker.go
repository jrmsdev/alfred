// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"sync"

	"github.com/jrmsdev/alfred/internal/core"
	"github.com/jrmsdev/alfred/internal/web"
	//~ "github.com/jrmsdev/alfred/log"
)

var Group = new(sync.WaitGroup)

func Start(name string) error {
	Group.Add(1)
	if name == "web" {
		return web.Start(Group)
	}
	return core.Start(Group)
}
