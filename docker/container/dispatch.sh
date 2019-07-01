#!/bin/bash
set -eu

SRCDIR=${ALFRED_SRC:-'/SRCDIR/notset'}
cd ${SRCDIR}

echo "-- dispatch"
echo "--     user $(id -a)"
echo "--     workdir $(pwd)"

go install ./devel/alfred-install
${GOPATH}/bin/alfred-install

/usr/local/bin/alfred-core -version
exec /usr/local/bin/alfred-core $@
