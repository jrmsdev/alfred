#!/usr/bin/env bash
set -eu
SRCDIR=${1:-'.'}
NAME='alfred'
IMAGE='alfred'
if test 'docs' = "${SRCDIR}"; then
	shift
	NAME='alfred-docs'
	IMAGE='alfred:docs'
elif test 'devel' = "${SRCDIR}"; then
	shift
	NAME='alfred-devel'
	IMAGE='alfred:dev'
fi
echo "-- run ${NAME}"
source ./docker/network.sh
docker run -it --rm --net=${NETNAME} --name=${NAME} \
	--add-host 'host.docker.internal:10.0.127.1' \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred ${IMAGE} $@
exit 0
