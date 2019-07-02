// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/exec"
	fpath "path/filepath"
)

var printf = fmt.Printf

const (
	VERSION = "19.7.1"
)

func main() {
	printf("alfred-install version %s\n", VERSION)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		printf("E: could not get GOPATH env var\n")
		os.Exit(1)
	}

	srcdir := fpath.Join(gopath, "src", "github.com", "jrmsdev", "alfred")
	err := os.Chdir(srcdir)
	if err != nil {
		printf("E: %s\n", err)
		os.Exit(1)
	}

	printf("  source dir %s\n", srcdir)
	cmdrun("./install.sh")

	printf("done\n")
}

func cmdrun(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	printf("  %s\n", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		printf("%s\n", err)
		printf("%s", out)
		os.Exit(2)
	}
}
