#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go assets cmd internal log
exit 0
