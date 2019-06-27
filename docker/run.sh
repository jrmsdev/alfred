#!/bin/sh -eu
docker run -it --rm --net=host --name=alfred \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred alfred $@
exit 0
