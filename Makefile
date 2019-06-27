GOPATH := /go
ALFRED_TEST ?=

.PHONY: build
build:
	go build -i ./...

.PHONY: clean
clean:
	go clean -i -cache -testcache ./...

.PHONY: check
check:
	go test $(ALFRED_TEST) ./...

.PHONY: gofmt
gofmt:
	gofmt -w -l -s _t lib
