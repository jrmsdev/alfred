#!/bin/bash
set -eu

SRCDIR=${ALFRED_SRC:-'/SRCDIR/notset'}
cd ${SRCDIR}

echo "-- dispatch"
echo "--     user $(id -a)"
echo "--     workdir $(pwd)"

go install ./cmd/alfred-install
${GOPATH}/bin/alfred-install

exec /usr/local/bin/alfred -version
