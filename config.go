// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package alfred

type alfredConfig struct {
	Log struct {
		Level string
	}
}

var Config = new(alfredConfig)
