// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	fpath "path/filepath"
)

var installPrefix = fpath.FromSlash("/usr/local")

func Prefix(args ...string) string {
	return fpath.Join(installPrefix, fpath.Join(args...))
}
