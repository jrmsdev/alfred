// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors

import (
	_errors "errors"
	"fmt"
)

type msg struct {
	Format string
}

var NotAbsPath = msg{"%s is not an absolute path: %s"}

func New(m msg, args ...interface{}) error {
	return _errors.New(fmt.Sprintf(m.Format, args...))
}
