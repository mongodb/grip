PACKAGES := ./ ./level ./message ./send

deps:
	go get github.com/coreos/go-systemd/journal

test-deps:deps
	go get gopkg.in/check.v1
	-go get github.com/alecthomas/gometalinter >/dev/null 2>&1
	-gometalinter --install --update >/dev/null 2>&1

build:deps
	go build -v

lint:
	gofmt -l $(PACKAGES)
	go vet $(PACKAGES)
	-gometalinter --disable=gotype --deadline=20s $(PACKAGES)

test:build lint
	go test -cover -v -check.v $(PACKAGES)
