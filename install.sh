#!/usr/bin/env bash
set -eu
DESTDIR=${DESTDIR:-''}
PREFIX=${PREFIX:-'/usr/local'}
GOBIN=${GOBIN:-"${DESTDIR}${PREFIX}/bin"}
#echo "-- GOBIN ${GOBIN}"
export GOBIN
for pkg in $(go list ./cmd/...); do
	cmd=$(basename ${pkg})
	echo "-- install ${GOBIN}/${cmd}"
	go install -i ${pkg}
done
exit 0
