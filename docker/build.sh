#!/usr/bin/env bash
set -eu
SRCDIR=${1:-'.'}
NAME='alfred'
if test 'docs' = "${SRCDIR}"; then
	NAME='alfred:docs'
elif test 'devel' = "${SRCDIR}"; then
	NAME='alfred:dev'
fi
echo "-- build ${NAME}"
docker build --rm --network host -t ${NAME} \
	--add-host 'host.docker.internal:127.0.0.1' ${SRCDIR}
exit 0
