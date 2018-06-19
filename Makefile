.PHONY: vendor build
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
	@dep ensure

update_lib:
	@dep ensure -update github.com/phrase/phraseapp-go

release:
	@sh build/build.sh
