// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package worker

import (
	"context"
	"os"
	"os/exec"
	fpath "path/filepath"

	"github.com/jrmsdev/alfred"
	"github.com/jrmsdev/alfred/log"
)

func Start(ctx context.Context, name string) error {
	bin := fpath.Join(alfred.Config.LibDir, "bin", name)
	log.Debug(bin)
	if _, err := os.Stat(bin); err != nil {
		log.Error(err)
		return err
	}
	cmd := exec.Command(bin)
	err := cmd.Run()
	if err != nil {
		log.Errorf("web worker: %s", err)
	}
	return err
}
