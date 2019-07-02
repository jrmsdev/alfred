#!/usr/bin/env bash
set -eu

DESTDIR=${DESTDIR:-''}
PREFIX=${PREFIX:-'/usr/local'}

GOBIN=${GOBIN:-"${DESTDIR}${PREFIX}/bin"}
export GOBIN

./gen.sh

for pkg in $(go list ./cmd/...); do
	src=${pkg/github\.com\/jrmsdev\/alfred\//}
	./gen.sh ${src}
	cmd=$(basename ${pkg})
	echo "-- install ${GOBIN}/${cmd}"
	go install -i ${pkg}
done

exit 0
