#!/usr/bin/env bash
set -eu

export BUILDDIR=${GOPATH}/pkg/alfred/build.run
./build.sh

exec ${BUILDDIR}/cmd/alfred -libdir ${BUILDDIR}/internal -log debug $@
