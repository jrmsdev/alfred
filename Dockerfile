FROM debian:testing-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.6.30"

USER root:root

RUN useradd -U -c alfred -d /home/alfred -m -s /bin/bash alfred

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

ENV GOPATH /go
ENV SRCDIR /go/src/github.com/jrmsdev/alfred

RUN mkdir -vp ${SRCDIR}

WORKDIR ${SRCDIR}
CMD /bin/bash ${SRCDIR}/docker/cmd.sh
