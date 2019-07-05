// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package web

import (
	"context"
	"os/exec"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/log"
)

func Start(ctx context.Context) {
	bin := fpath.Join(alfred.Config.LibDir, "bin", "web")
	log.Debug(bin)
	cmd := exec.Command(bin)
	err := cmd.Run()
	if err != nil {
		log.Errorf("web worker: %s", err)
	}
}
