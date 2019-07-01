FROM debian:testing-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.7.1"

USER root:root

ARG ALFRED_UID
ARG ALFRED_GID

RUN groupadd -g ${ALFRED_GID} alfred
RUN useradd -c alfred -m -d /home/alfred -g alfred -s /bin/bash -u ${ALFRED_UID} alfred

ENV DEBIAN_FRONTEND noninteractive

COPY docker/apt-proxy.conf /etc/apt/apt.conf.d/02proxy

RUN apt-get clean
RUN apt-get update
RUN apt-get dist-upgrade -yy --purge

RUN apt-get install -yy --no-install-recommends golang

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

USER alfred:alfred

ENV GOPATH /home/alfred/go
ENV ALFRED_SRC ${GOPATH}/src/github.com/jrmsdev/alfred

RUN mkdir -vp ${ALFRED_SRC}
WORKDIR ${ALFRED_SRC}

RUN go version
RUN go env | sort

ENTRYPOINT ["./docker/container/dispatch.sh"]
