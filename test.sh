#!/usr/bin/env bash
set -eu
ALFRED_TEST=${ALFRED_TEST:-''}
go test ${ALFRED_TEST} ./...
exit 0
