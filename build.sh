#!/usr/bin/env bash
set -eu

PKGLIST="$@"
BUILDDIR=${BUILDDIR:-"${GOPATH}/pkg/alfred/build"}

mkdir -p ${BUILDDIR}
if test "X" = "X${PKGLIST}"; then
	PKGLIST=$(cat ./build.pkg)
fi

./gen.sh

echo "-- build dir ${BUILDDIR}"
for pkg in ${PKGLIST}; do
	dst=${BUILDDIR}/${pkg}
	echo "-- build ${pkg}"
	./gen.sh ${pkg}
	go build -i -o ${dst} ./${pkg}
done

exit 0
