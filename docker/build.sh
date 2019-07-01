#!/usr/bin/env bash
set -eu
SRCDIR=${1:-'.'}
NAME='alfred'
if test 'docs' = "${SRCDIR}"; then
	NAME='alfred:docs'
elif test 'devel' = "${SRCDIR}"; then
	NAME='alfred:dev'
fi
ALFRED_UID=$(id -u)
ALFRED_GID=$(id -g)
echo "-- build ${NAME}"
docker build --rm --network host -t jrmsdev/${NAME} \
	--build-arg ALFRED_UID=$(id -u) \
	--build-arg ALFRED_GID=$(id -g) \
	--add-host 'host.docker.internal:127.0.0.1' \
	${SRCDIR}
exit 0
