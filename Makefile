.PHONY: vendor build dist

default: build

all: build test vet

test:
	@go test ./...

vet:
	@go vet ./...

build:
	@go get ./...

vendor:
	@dep ensure

update_lib:
	@dep ensure -update github.com/phrase/phraseapp-go

release:
	sh build/release.sh

dist:
	sh build/build.sh
	sh build/innosetup/create_installer.sh
