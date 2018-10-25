export CGO_ENABLED:=0

VERSION=$(shell ./scripts/git-version)
GOPATH_BIN:=$(shell echo ${GOPATH} | awk 'BEGIN { FS = ":" }; { print $1 }')/bin

.PHONY: all
all: build

.PHONY: build
build:
	@go build -o bin/demoapp-pcf-healthcheck -v github.com/alekssaul/demoapp-pcf-healthcheck

.PHONY: vendor
vendor:
	@glide update --strip-vendor

.PHONY: test
test:
	@go test ./