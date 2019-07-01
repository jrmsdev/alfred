FROM jrmsdev/alfred

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.7.1"

USER root:root

ARG ALFRED_UID
ARG ALFRED_GID

RUN apt-get clean
RUN apt-get update

RUN apt-get install -yy --no-install-recommends sudo less golang-golang-x-tools\
	go-dep gcc libc-dev

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

RUN echo 'alfred ALL=NOPASSWD: ALL' >/etc/sudoers.d/alfred
RUN chmod 440 /etc/sudoers.d/alfred

ENV PAGER less

USER alfred:alfred
CMD exec /bin/bash -i -l
