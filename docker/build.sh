#!/usr/bin/env bash
set -eu
SRCDIR=${1:-'.'}
NAME='alfred'
if test 'docs' = "${SRCDIR}"; then
	NAME='alfredocs'
fi
docker build -t ${NAME} --network host ${SRCDIR}
exit 0
