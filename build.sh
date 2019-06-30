#!/usr/bin/env bash
set -eu
BUILDDIR=${GOPATH}/pkg/alfred/build
mkdir -p ${BUILDDIR}
for pkg in $(go list ./cmd/...); do
	cmd=${BUILDDIR}/$(basename ${pkg})
	echo "-- build ${cmd}"
	go build -i -o ${cmd} ${pkg}
done
exit 0
