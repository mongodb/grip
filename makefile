PACKAGES := ./ ./level ./message ./send

deps:
	go get -u github.com/coreos/go-systemd/journal

test-deps:
	go get -u gopkg.in/check.v1

build:deps
	go build -v

lint:
	gofmt -l $(PACKAGES)
	go vet $(PACKAGES)

test:build lint
	go test -v -check.v $(PACKAGES)
