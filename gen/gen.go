// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Template(name string, data *map[string]string) {
	fn := filepath.FromSlash(name)

	dstfn := fn
	if strings.HasSuffix(name, ".in") {
		dstfn = strings.Replace(dstfn, ".in", "", 1)
	}
	fmt.Printf("-- generate %s\n", dstfn)

	tpl := template.New(name)

	src, err := os.Open(fn)
	check(err)
	defer src.Close()

	if blob, err := ioutil.ReadAll(src); err != nil {
		panic(err)
	} else {
		tpl, err = tpl.Parse(string(blob))
		check(err)
	}

	dst, err2 := os.Create(dstfn)
	check(err2)
	defer dst.Close()

	err = tpl.Execute(dst, data)
	check(err)
}
