PACKAGES := ./ ./level ./message ./send

deps:
	go get github.com/coreos/go-systemd/journal

test-deps:deps
	go get gopkg.in/check.v1

build:deps
	go build -v

lint:
	gofmt -l $(PACKAGES)
	go vet $(PACKAGES)

test:build lint
	go test -cover -v -check.v $(PACKAGES)
