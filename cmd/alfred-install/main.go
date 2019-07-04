// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	fpath "path/filepath"
	"strings"
)

var printf = fmt.Printf

func getversion(srcdir string) string {
	fh, err := os.Open(fpath.Join(srcdir, "VERSION"))
	if err != nil {
		printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	defer fh.Close()
	r := bufio.NewScanner(fh)
	r.Scan()
	return strings.TrimSpace(r.Text())
}

func cmdrun(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	printf("  %s\n", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		printf("ERROR: %s\n", err)
		printf("%s", out)
		os.Exit(2)
	}
}

func main() {
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

	printf("alfred-install version %s\n", getversion(srcdir))
	printf("  source dir %s\n", srcdir)

	cmdrun("./install.sh")
	printf("done\n")
}
