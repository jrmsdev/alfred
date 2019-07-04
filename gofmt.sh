#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go *.go.in
gofmt -w -l -s assets cmd errors internal log
exit 0
