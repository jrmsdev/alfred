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
	gofmt -w -l -s internal lib

.PHONY: goenv
goenv:
	@go env | sort -t '=' -k1,1
