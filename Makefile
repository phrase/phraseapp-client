.PHONY: vendor
PACKAGES = $(shell go list ./... | grep -v "/vendor")

default: build

all: build test vet

test:
	@go test ${PACKAGES}

vet:
	@go vet ${PACKAGES}

build:
	@go get ${PACKAGES}

vendor:
	godep save ./...
	@godep save ./...

update_lib:
	godep update github.com/phrase/phraseapp-go/...
	make vendor
