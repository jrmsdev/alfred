#!/bin/sh -eu
docker run -it --rm --net=host --name=alfredocs \
	-v ${PWD}:/go/src/github.com/jrmsdev/alfred/docs alfredocs $@
exit 0
