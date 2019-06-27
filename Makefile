GOPATH := /go

.PHONY: build
build:
	go build -i ./...

.PHONY: clean
clean:
	go clean -i -cache -testcache ./...

.PHONY: check
check:
	go test ./...

.PHONY: gofmt
gofmt:
	gofmt -w -l -s _t lib
