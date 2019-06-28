#!/usr/bin/env bash
set -eu
gofmt -w -l -s bin internal lib
exit 0
