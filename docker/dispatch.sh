#!/bin/bash
set -eu

SRCDIR=${SRCDIR:-'/SRCDIR/notset'}
cd ${SRCDIR}

umask ${ALFRED_UMASK}

echo "-- dispatch"
echo "--     user $(id -a)"
echo "--     umask $(umask)"
echo "--     workdir $(pwd)"

./install.sh
exec /usr/local/bin/alfred $@
