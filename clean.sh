#!/usr/bin/env bash
set -eu
BUILDDIR=${GOPATH}/pkg/alfred/build
go clean -i -testcache ./...
rm -vrf ${BUILDDIR} ${BUILDDIR}.*
rm -vf coverage.out coverage.html \
	build_info.go \
	version_info.go
exit 0
