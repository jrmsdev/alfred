#!/usr/bin/env bash
set -eu
if which dep >/dev/null; then
	dep check
fi
go vet ./...
./test.sh
exit 0
