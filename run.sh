#!/usr/bin/env bash
set -eu

CMD=${1:?'command pkg path?'}
export BUILDDIR=${GOPATH}/pkg/alfred/build.run

echo ${CMD} | grep -E '^cmd\/|^internal\/bin' >/dev/null

./build.sh ${CMD}

shift
#echo "-- exec ${BUILDDIR}/${CMD} $@"
exec ${BUILDDIR}/${CMD} $@
