FROM jrmsdev/alfred

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="19.7.1"

USER root:root

ARG ALFRED_UID
ARG ALFRED_GID

RUN apt-get clean
RUN apt-get update

RUN apt-get install -yy --no-install-recommends sudo procps gnupg2 curl git \
	ca-certificates

RUN echo 'alfred ALL=NOPASSWD: ALL' >/etc/sudoers.d/alfred
RUN chmod 440 /etc/sudoers.d/alfred

USER alfred:alfred

RUN gpg2 --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB

RUN curl -sSL -o /var/tmp/rvm.setup https://get.rvm.io
RUN chmod 750 /var/tmp/rvm.setup
RUN /var/tmp/rvm.setup stable --ruby

RUN bash -c 'source ~/.rvm/scripts/rvm && ruby --version'

COPY --chown=alfred:alfred ./Gemfile docs/
COPY --chown=alfred:alfred ./Gemfile.lock docs/

WORKDIR ${ALFRED_SRC}/docs
RUN bash -c 'source ~/.rvm/scripts/rvm && bundle install'

WORKDIR ${ALFRED_SRC}
USER root:root

RUN apt-get clean
RUN apt-get autoremove -yy --purge

RUN rm -rf /var/lib/apt/lists/*
RUN rm -f /var/cache/apt/archives/*.deb
RUN rm -f /var/cache/apt/*cache.bin

USER alfred:alfred
CMD ./docker/container/docsdevel.sh
