#!/usr/bin/env bash
set -eu

CMD=${1:?'command pkg path?'}
BUILDDIR=${GOPATH}/pkg/alfred/build

echo ${CMD} | grep -E '^\.\/cmd\/alfred' >/dev/null

./build.sh ${CMD}

shift
echo "-- exec ${BUILDDIR}/${CMD} $@"
exec ${BUILDDIR}/${CMD} $@
