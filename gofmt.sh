#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go *.go.in \
	assets \
	cmd \
	devel \
	gen \
	internal \
	log
exit 0
