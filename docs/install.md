# Install

Requires Go version >=1.11.6

Use `alfred-install` command to build and install from source under current
user's home directory:

	PREFIX=${HOME}
	GOPATH=${PREFIX}
	PATH=${GOPATH}/bin:${PATH}

	go get -u github.com/jrmsdev/alfred
	go install github.com/jrmsdev/alfred/cmd/alfred-install

	alfred-install

If PREFIX is not set, `alfred` is installed under `/usr/local` by default.
DESTDIR env var is also honored if set.

	DESTDIR=/chroot/alfred alfred-install

Or explicitly set destination directories:

	ALFRED_BINDIR=/usr/bin
	ALFRED_LIBDIR=/usr/lib/alfred
	alfred-install
