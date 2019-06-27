#!/bin/sh -eu
docker build -t alfred --network host .
exit 0
