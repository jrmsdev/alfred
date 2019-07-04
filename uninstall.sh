#!/usr/bin/env bash
set -eu

DESTDIR=${DESTDIR:-''}
PREFIX=${PREFIX:-'/usr/local'}

BINDIR=${DESTDIR}${PREFIX}/bin
LIBDIR=${DESTDIR}${PREFIX}/lib/alfred

for pkg in $(cat build.pkg); do
	dstdir=${BINDIR}
	if echo ${pkg} | grep -E '^cmd\/' >/dev/null; then
		dst=${dstdir}/$(basename ${pkg})
		if test -e ${dst}; then
			rm -vf ${dst}
		fi
	fi
done

if test -d ${LIBDIR}; then
	rm -vrf ${LIBDIR}
fi

exit 0
