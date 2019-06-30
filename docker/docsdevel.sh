#!/bin/bash
set -eu
SRCDIR=${SRCDIR:-'/SRCDIR/notset'}
echo "-- user $(id -a)"
echo "-- umask $(umask)"
cd ${SRCDIR}
echo "-- workdir $(pwd)"
source ~/.rvm/scripts/rvm
jekyll serve -H 0.0.0.0 -P 8080
