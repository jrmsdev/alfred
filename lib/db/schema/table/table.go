// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package table

type Table struct {
	name string
}

func New(name string) *Table {
	return &Table{name}
}
