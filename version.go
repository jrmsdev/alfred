// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

import (
	"fmt"
)

var version = "master"
var version_build = "devel"

func GetVersion() string {
	return version
}

func Version() string {
	return fmt.Sprintf("version %s", version)
}

func GetVersionBuild() string {
	return version_build
}

func VersionBuild() string {
	return fmt.Sprintf("build %s", version_build)
}
