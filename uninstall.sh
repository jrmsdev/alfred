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
			echo "-- remove ${dst}"
			rm -f ${dst}
		fi
	fi
done

if test -d ${LIBDIR}; then
	echo "-- remove ${LIBDIR}"
	rm -rf ${LIBDIR}
fi

exit 0
