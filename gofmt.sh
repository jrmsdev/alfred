#!/usr/bin/env bash
set -eu
gofmt -w -l -s *.go *.go.in \
	assets \
	cmd \
	devel \
	internal \
	libexec \
	log
exit 0
