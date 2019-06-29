FROM debian:testing-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.6.27"

RUN useradd -U -c alfred -d /home/alfred -m -s /bin/bash alfred

ENV DEBIAN_FRONTEND noninteractive

COPY docker/apt-proxy.conf /etc/apt/apt.conf.d/02proxy

RUN apt-get clean
RUN apt-get update

RUN apt-get dist-upgrade -yy --purge
RUN apt-get install -yy --no-install-recommends sudo less golang \
	golang-golang-x-tools

RUN apt-get install -yy --no-install-recommends gnupg2 curl ca-certificates \
	procps git

RUN echo "PS1='docker:\W# '" >>/root/.bashrc

COPY docker/sudoers /etc/sudoers.d/alfred
RUN chmod 440 /etc/sudoers.d/alfred

ENV GOPATH /go
RUN mkdir /go
RUN chown alfred:alfred /go
RUN chmod 750 /go

USER alfred:alfred

RUN gpg2 --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB
RUN curl -sSL https://get.rvm.io | bash -s stable --ruby

RUN mkdir /var/tmp/docs
COPY docs/Gemfile /var/tmp/docs/Gemfile
COPY docs/Gemfile.lock /var/tmp/docs/Gemfile.lock
RUN (cd /var/tmp/docs && bundle install)
RUN rm -rf /var/tmp/docs

USER root:root

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

USER alfred:alfred

RUN echo '### alfred setup' >>/home/alfred/.bashrc
RUN echo 'umask 027' >>/home/alfred/.bashrc
RUN echo "PS1='docker:\W\$ '" >>/home/alfred/.bashrc

RUN (umask 027 && mkdir -p /go/src/github.com/jrmsdev/alfred)
WORKDIR /go/src/github.com/jrmsdev/alfred

CMD /bin/bash
