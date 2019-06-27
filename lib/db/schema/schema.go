// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"container/list"

	"github.com/jrmsdev/alfred/lib/db/schema/table"
)

type Schema struct {
	stmt *list.List
}

func New() *Schema {
	return &Schema{stmt: list.New()}
}

func (s *Schema) CreateTable(name string) (*table.Table, error) {
	tbl := table.New(name)
	return tbl, nil
}
