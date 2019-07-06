// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package model

type Status struct {
	Id int
	Status string `orm:"size(50)"`
}

func GetStatus() string {
	return "init"
}
