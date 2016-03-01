default: build

all: build test vet

test:
	go test ./...

vet:
	go vet ./...

build:
	go get ./...
