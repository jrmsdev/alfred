#!/usr/bin/env bash
set -eu
# use GOPATH/pkg for development
binfn=${GOPATH}/pkg/alfred-install
rm -f ${binfn}
go build -i -o ${binfn} ./devel/alfred-install
exec ${binfn}
