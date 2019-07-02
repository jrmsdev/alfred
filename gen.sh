#!/usr/bin/env bash
set -eu
PKG=${1:-''}
if test "X" != "X${PKG}"; then
	PKG=${PKG}/
fi
genfile="${PKG}generate.go"
if test -s ${genfile}; then
	echo "-- ${genfile}"
	go generate ${genfile}
fi
exit 0
