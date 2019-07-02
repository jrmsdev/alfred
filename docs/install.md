# Install

Requires Go version >=1.11.6

	$ GOPATH=~/go
	$ PATH=${GOPATH}/bin:${PATH}

	$ go get -u github.com/jrmsdev/alfred
	$ go install github.com/jrmsdev/alfred/cmd/alfred-install

	$ alfred-install

DESTDIR and PREFIX env vars are honored

	$ DESTDIR=/chroot/alfred PREFIX=${HOME} alfred-install
