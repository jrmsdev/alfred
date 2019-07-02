#!/usr/bin/env bash
set -eu
gols() {
	go list ./... | sed 's#github.com/jrmsdev/alfred/##' | tail -n +2
}
gofmt -w -l -s *.go *.go.in
gofmt -w -l -s $(gols)
exit 0
