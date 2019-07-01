#!/usr/bin/env bash
set -eu
PKG=${1:-'.'}
genfile=${PKG}/generate.go
if test -s ${genfile}; then
	echo "-- ${genfile}"
	go generate ${genfile}
fi
exit 0
