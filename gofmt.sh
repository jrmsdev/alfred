#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go *.go.in \
	assets \
	cmd \
	gen \
	internal \
	log
exit 0
