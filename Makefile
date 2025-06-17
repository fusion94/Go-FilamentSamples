.PHONY: build clean test lint fmt vet install run help

BINARY_NAME=filament-samples
MAIN_PATH=./cmd/filament-samples
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  install  - Install the binary to GOPATH/bin"
	@echo "  run      - Run the application"
	@echo "  test     - Run tests"
	@echo "  lint     - Run golangci-lint"
	@echo "  fmt      - Format code"
	@echo "  vet      - Run go vet"
	@echo "  clean    - Clean build artifacts"

build:
	go mod tidy
	go build $(LDFLAGS) -o $(BINARY_NAME) -v $(MAIN_PATH)

install:
	go mod tidy
	go install $(LDFLAGS) -v $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

test:
	go test -v ./...

lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not found, install it from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY_NAME)
	go clean ./..
