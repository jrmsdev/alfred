#!/usr/bin/env bash
set -eu
if which dep >/dev/null; then
	dep check
fi
./gofmt.sh
go vet ./...
./test.sh
exit 0
