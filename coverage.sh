#!/usr/bin/env bash
set -eu
go test -coverprofile coverage.out ./...
go tool cover -html coverage.out -o coverage.html
exit 0
