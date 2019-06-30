#!/bin/bash
set -eu

SRCDIR=${SRCDIR:-'/SRCDIR/notset'}
cd ${SRCDIR}

umask ${ALFRED_UMASK}

echo "-- login"
echo "--     user $(id -a)"
echo "--     umask $(umask)"
echo "--     workdir $(pwd)"

exec /bin/bash
