// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	"fmt"
)

var version = "master"

func Version(prog string) string {
	return fmt.Sprintf("%s version %s", prog, version)
}
