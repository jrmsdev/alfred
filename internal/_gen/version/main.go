// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jrmsdev/alfred/internal/_gen"
)

var data = map[string]string{
	"VMajor": "0",
	"VMinor": "0",
	"VPatch": "0",
}

func getversion() string {
	fh, err := os.Open("VERSION")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	defer fh.Close()
	r := bufio.NewScanner(fh)
	r.Scan()
	return strings.TrimSpace(r.Text())
}

func main() {
	v := getversion()
	args := strings.Split(v, ".")
	data["VMajor"] = args[0]
	data["VMinor"] = args[1]
	if len(args) == 3 {
		data["VPatch"] = args[2]
	}
	gen.Template("version_info.go.in", &data)
}
