#!/bin/bash
set -eu
SRCDIR=${SRCDIR:-'/SRCDIR/notset'}
echo "-- user $(id -a)"
echo "-- umask $(umask)"
cd ${SRCDIR}
echo "-- workdir $(pwd)"
./install.sh
exec /usr/local/bin/alfred $@
