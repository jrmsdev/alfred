#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go *.go.in \
	assets \
	cmd \
	devel \
	internal \
	log
exit 0
