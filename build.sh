#!/usr/bin/env bash
set -eu

PKGLIST="$@"
BUILDDIR=${GOPATH}/pkg/alfred/build

mkdir -p ${BUILDDIR}
if test "X" = "X${PKGLIST}"; then
	PKGLIST='./cmd/...'
fi

go generate generate.go

for pkg in $(go list ${PKGLIST}); do
	dst=${pkg/github\.com\/jrmsdev\/alfred\//}
	ext='.a'
	if echo ${dst} | grep -E '^cmd\/' >/dev/null; then
		ext=''
	fi
	dst=${BUILDDIR}/${dst}${ext}
	echo "-- build ${dst}"
	go build -i -o ${dst} ${pkg}
done

exit 0
