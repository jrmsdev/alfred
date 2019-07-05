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
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	} else {
		args = make([]string, 0)
	}
	bin := fpath.Join(alfred.Config.LibDir, "bin", name)
	if _, err := os.Stat(bin); err != nil {
		log.Error(err)
		return err
	}
	log.Debug("%s %v", bin, args)
	cmd := exec.CommandContext(ctx, bin, args...)
	err := cmd.Start()
	if err != nil {
		log.Errorf("%s: %s", name, err)
	}
	log.Debug("%d %s", cmd.Process.Pid, cmd.Path)
	return cmd.Wait()
}
