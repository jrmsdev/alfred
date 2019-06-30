#!/bin/bash
set -eu

DOCKER_CMD=${1:-'dispatch'}
ALFRED_UID=${ALFRED_UID:-'1000'}
ALFRED_GID=${ALFRED_GID:-'1000'}
ALFRED_UMASK=${ALFRED_UMASK:-'0022'}
GOPATH=${GOPATH:-'/go'}
SRCDIR=${SRCDIR:-`pwd`}

test '0' = $(id -u)

echo "-- cmd ${DOCKER_CMD}"
echo "--     uid ${ALFRED_UID}"
echo "--     gid ${ALFRED_GID}"
echo "--     umask ${ALFRED_UMASK}"

echo "umask ${ALFRED_UMASK}" >>/home/alfred/.bashrc
echo "export GOPATH=${GOPATH}" >>/home/alfred/.bashrc
echo "export SRCDIR=${SRCDIR}" >>/home/alfred/.bashrc

usermod -u ${ALFRED_UID} -g ${ALFRED_GID} alfred

chown alfred:alfred ${GOPATH}
chgrp alfred /usr/local/bin
chmod 0775 /usr/local/bin

su -c "/bin/bash -l ${SRCDIR}/docker/${DOCKER_CMD}.sh" alfred
exit 0
