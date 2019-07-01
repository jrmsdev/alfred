#!/bin/bash

SRCDIR=${SRCDIR:-'/SRCDIR/notset'}
cd ${SRCDIR}/docs

umask ${ALFRED_UMASK}

echo "-- docsdevel"
echo "--     user $(id -a)"
echo "--     umask $(umask)"
echo "--     workdir $(pwd)"

source ~/.rvm/scripts/rvm
exec jekyll serve -H 0.0.0.0 -P 8080
