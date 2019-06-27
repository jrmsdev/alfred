// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"testing"

	"github.com/jrmsdev/alfred/_t/check"
)

func TestSchema(t *testing.T) {
	s := New()
	if check.NotEqual(t, s.stmt.Len(), 0, "schema stmt should be empty") {
		t.FailNow()
	}
	s.CreateTable("testing")
}
