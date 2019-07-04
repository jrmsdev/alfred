#!/usr/bin/env bash
set -eu

PKGLIST="$@"
BUILDDIR=${BUILDDIR:-"${GOPATH}/pkg/alfred/build"}

mkdir -p ${BUILDDIR}
if test "X" = "X${PKGLIST}"; then
	PKGLIST=$(cat ./build.pkg)
fi

export ALFRED_BINDIR=${ALFRED_BINDIR:-'/usr/local/bin'}
export ALFRED_LIBDIR=${ALFRED_LIBDIR:-'/usr/local/lib/alfred'}

./gen.sh

echo "-- build dir ${BUILDDIR}"
for pkg in ${PKGLIST}; do
	dst=${BUILDDIR}/${pkg}
	echo "-- build ${pkg}"
	./gen.sh ${pkg}
	go build -i -o ${dst} ./${pkg}
done

exit 0
