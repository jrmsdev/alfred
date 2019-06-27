FROM debian:testing-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.6.27"

RUN useradd -U -c alfred -d /home/alfred -m -s /bin/bash alfred

ENV DEBIAN_FRONTEND noninteractive

COPY docker/apt-proxy.conf /etc/apt/apt.conf.d/02proxy

RUN apt-get clean
RUN apt-get update

RUN apt-get dist-upgrade -yy --purge
RUN apt-get install -yy --no-install-recommends sudo golang \
	golang-golang-x-tools make less

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

RUN echo "PS1='docker:\W# '" >>/root/.bashrc

COPY docker/sudoers /etc/sudoers.d/alfred
RUN chmod 440 /etc/sudoers.d/alfred

ENV GOPATH /go
RUN mkdir /go
RUN chown alfred:alfred /go
RUN chmod 750 /go

USER alfred:alfred

RUN echo '### alfred setup' >>/home/alfred/.bashrc
RUN echo 'umask 027' >>/home/alfred/.bashrc
RUN echo "PS1='docker:\W\$ '" >>/home/alfred/.bashrc

RUN (umask 027 && mkdir -p /go/src/github.com/jrmsdev/alfred)
WORKDIR /go/src/github.com/jrmsdev/alfred

CMD /bin/bash
