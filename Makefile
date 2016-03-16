default: build

all: build test vet

test:
	go test ./...

vet:
	go vet ./...

build:
	go get ./...

godep:
	godep save -r ./...
	@godep save -r ./...

update_lib:
	godep update github.com/phrase/phraseapp-go/...
	make godep
