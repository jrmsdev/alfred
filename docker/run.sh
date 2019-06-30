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
docker run -it --rm --network ${NETNAME} --name ${NAME} --hostname ${NAME} \
	--add-host 'host.docker.internal:10.0.127.1' \
	--network-alias "${NAME}.docker.internal" \
	-p 127.0.0.1:8080:8080 \
	-p 127.0.0.1:8180:8180 \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred ${IMAGE} $@
exit 0
