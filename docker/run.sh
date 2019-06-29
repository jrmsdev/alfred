#!/usr/bin/env bash
set -eu
SRCDIR=${1:-'.'}
NAME='alfred'
if test 'docs' = "${SRCDIR}"; then
	shift
	NAME='alfredocs'
fi
docker run -it --rm --net=host --name=${NAME} \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred ${NAME} $@
exit 0
