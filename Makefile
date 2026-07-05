.PHONY: all build test vet integration clean coverage

all: fmt vet lint test build

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

integration:
	go run ./examples/integration

clean:
	go clean ./...
	@rm -f coverage.out coverage.html integration

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

fmt:
	gofmt -w -s .

lint:
	staticcheck ./...

precommit: fmt lint vet test build
