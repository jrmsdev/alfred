#!/usr/bin/env bash
set -eu

mkdir -p ${GOPATH}/pkg/alfred
export BUILDDIR=$(mktemp -d ${GOPATH}/pkg/alfred/build.XXXXXXX)

./build.sh

DESTDIR=${DESTDIR:-''}
PREFIX=${PREFIX:-'/usr/local'}

BINDIR=${DESTDIR}${PREFIX}/bin
LIBDIR=${DESTDIR}${PREFIX}/lib/alfred

mkdir -p -m 0755 ${BINDIR} ${LIBDIR}

for pkg in $(cat build.pkg); do
	src=${BUILDDIR}/${pkg}
	if ! test -s ${src}; then
		echo "E: ${src} file not found, build failed!?" >&2
	fi
	dstdir=${LIBDIR}
	if echo ${pkg} | grep -E '^cmd\/' >/dev/null; then
		dstdir=${BINDIR}
	fi
	dst=${dstdir}/$(basename ${pkg})
	echo "-- install ${dst}"
	install -m 0755 ${src} ${dst}
done

rm -rf ${BUILDDIR}
exit 0
