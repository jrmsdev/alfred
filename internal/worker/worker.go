// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"sync"

	//~ "github.com/jrmsdev/alfred/log"
)

type Group struct {
	sync.WaitGroup
}
