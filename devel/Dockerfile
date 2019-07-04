FROM jrmsdev/alfred

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.7.3"

USER root:root

ARG ALFRED_UID
ARG ALFRED_GID

RUN apt-get clean
RUN apt-get update

RUN apt-get install -yy --no-install-recommends sudo less golang-golang-x-tools\
	gcc libc-dev

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

RUN echo 'alfred ALL=NOPASSWD: ALL' >/etc/sudoers.d/alfred
RUN chmod 440 /etc/sudoers.d/alfred

USER alfred:alfred

ENV PAGER less
ENV ALFRED_CORE 'alfred-devel:27719'
ENV ALFRED_WEB 'alfred-devel:21680'

CMD exec /bin/bash -i -l
