#!/usr/bin/env bash
set -eu
BUILDDIR=${GOPATH}/pkg/alfred/build
go clean -i -cache -testcache ./...
rm -f coverage.out coverage.html \
	version_info.go
rm -rf ${BUILDDIR}
exit 0
