FROM alfred

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.6.28"

USER root:root

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update

RUN apt-get install -yy --no-install-recommends gnupg2 curl ca-certificates \
	procps git

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

USER alfred:alfred

RUN gpg2 --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB

RUN curl -sSL -o /var/tmp/rvm.setup https://get.rvm.io
RUN chmod 750 /var/tmp/rvm.setup
RUN /var/tmp/rvm.setup stable --ruby

RUN bash -c 'source ~/.rvm/scripts/rvm && ruby --version'
RUN echo 'source ~/.rvm/scripts/rvm' >>/home/alfred/.bashrc

COPY ./Gemfile docs/
COPY ./Gemfile.lock docs/
RUN sudo chown -R alfred:alfred docs

WORKDIR ${SRCDIR}/docs
RUN bash -c 'source ~/.rvm/scripts/rvm && bundle install'

CMD /bin/bash -c 'source ~/.rvm/scripts/rvm && jekyll serve'
