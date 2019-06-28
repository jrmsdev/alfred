// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/alfred/lib/log"
)

func main() {
	log.Init("debug")
	log.Debug("alfred main")
	log.Warnf("warning")
	log.Errorf("error")
	log.Print("msg")
	log.Debug("end")
}
