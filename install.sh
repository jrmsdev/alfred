#!/usr/bin/env bash
set -eu
GOBIN=${GOBIN:-''}
GOEXE=${GOEXE:-''}
export GOBIN
export GOEXE
for pkg in $(go list ./cmd/...); do
	cmd=$(basename ${pkg})
	echo "-- install ${cmd}"
	go install -i ${pkg}
done
exit 0
