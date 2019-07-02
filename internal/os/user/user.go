// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package user

import (
	"os"
	"os/user"
	fpath "path/filepath"
)

type osUser struct {
	Uid  string
	Gid  string
	Name string
	Home string
}

var cur *osUser

func init() {
	cur = new(osUser)
	if real, err := user.Current(); err != nil {
		cur.Uid = "65534"
		cur.Gid = "65534"
		cur.Name = "nobody"
	} else {
		cur.Uid = real.Uid
		cur.Gid = real.Gid
		cur.Name = real.Username
	}
	if cur.Home == "" {
		cur.Home = os.Getenv("HOME")
	}
	if cur.Home == "" {
		cur.Home = fpath.FromSlash("/empty/home")
	}
}

func Uid() string {
	return cur.Uid
}

func Gid() string {
	return cur.Gid
}

func Name() string {
	return cur.Name
}

func Home(args ...string) string {
	return fpath.Join(cur.Home, fpath.Join(args...))
}
