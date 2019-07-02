#!/usr/bin/env bash
set -eu

SRCDIR=${1:-'.'}

NAME='alfred'
IMAGE='alfred'
PUBLISH='-p 127.0.0.1:21600:21600 -p 127.0.0.1:21680:21680'

if test 'docs' = "${SRCDIR}"; then
	NAME='alfred-docs'
	IMAGE='alfred:docs'
	PUBLISH='-p 127.0.0.1:8080:8080'
elif test 'devel' = "${SRCDIR}"; then
	NAME='alfred-devel'
	IMAGE='alfred:dev'
fi

if test 'alfred' != ${NAME}; then
	shift
fi

source ./docker/networkrc

echo "-- run ${NAME}"
docker run -it --rm --network ${NETNAME} --name ${NAME} --hostname ${NAME} \
	${PUBLISH} \
	--add-host 'host.docker.internal:10.0.127.1' \
	--network-alias "${NAME}.docker.internal" \
	-v ${PWD}:/home/alfred/go/src/github.com/jrmsdev/alfred \
	jrmsdev/${IMAGE} $@

exit 0
