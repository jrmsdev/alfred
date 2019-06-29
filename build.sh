#!/usr/bin/env bash
set -eu
mkdir -p ./build
for pkg in $(go list ./cmd/...); do
	cmd=$(basename ${pkg})
	go build -i -o ./build/${cmd} ${pkg}
done
exit 0
