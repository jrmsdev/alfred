#!/usr/bin/env bash
set -eu
go clean -i -cache -testcache ./...
rm -f coverage.out coverage.html
exit 0
