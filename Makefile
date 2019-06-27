GOPATH ?= /go
ALFRED_TEST ?=

.PHONY: build
build:
	go build -i ./...

.PHONY: clean
clean:
	go clean -i -cache -testcache ./...
	rm -f coverage.out coverage.html

.PHONY: check
check:
	go test $(ALFRED_TEST) ./...

.PHONY: coverage
coverage:
	go test -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

.PHONY: gofmt
gofmt:
	gofmt -w -l -s internal lib

.PHONY: goenv
goenv:
	@go env | sort -t '=' -k1,1
