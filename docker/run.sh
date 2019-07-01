#!/usr/bin/env bash
set -eu

SRCDIR=${1:-'.'}
DOCKER_CMD=${2:-'dispatch'}

NAME='alfred'
IMAGE='alfred'

if test 'docs' = "${SRCDIR}"; then
	NAME='alfred-docs'
	IMAGE='alfred:docs'
elif test 'devel' = "${SRCDIR}"; then
	NAME='alfred-devel'
	IMAGE='alfred:dev'
fi

ALFRED_UID=$(id -u)
ALFRED_GID=$(id -g)
ALFRED_UMASK=$(umask)
echo "-- run ${NAME}"
echo "--     uid ${ALFRED_UID}"
echo "--     gid ${ALFRED_GID}"
echo "--     umask ${ALFRED_UMASK}"

source ./docker/network.sh

docker run -it --rm --network ${NETNAME} --name ${NAME} --hostname ${NAME} \
	-e ALFRED_UID=${ALFRED_UID} \
	-e ALFRED_GID=${ALFRED_GID} \
	-e ALFRED_UMASK=${ALFRED_UMASK} \
	--add-host 'host.docker.internal:10.0.127.1' \
	--network-alias "${NAME}.docker.internal" \
	-p 127.0.0.1:8080:8080 \
	-p 127.0.0.1:8180:8180 \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred \
	jrmsdev/${IMAGE} /bin/bash ./docker/cmd.sh ${DOCKER_CMD}

exit 0
