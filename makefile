PACKAGES := ./ ./level ./message ./send
projectPath := github.com/tychoish/grip
coverageFile := coverage.out

deps:
	go get github.com/coreos/go-systemd/journal

test-deps:deps
	go get github.com/stretchr/testify
	-go get github.com/alecthomas/gometalinter >/dev/null 2>&1
	-gometalinter --install --update >/dev/null 2>&1

build:deps
	go build -v

lint:
	gofmt -l $(PACKAGES)
	go vet $(PACKAGES)
	-gometalinter --disable=gotype --deadline=20s $(PACKAGES)

test:build
	go test -v -covermode=count -coverprofile=${coverageFile} ${projectPath}
	go tool cover -func=${coverageFile}

coverage-report:test
	go tool cover -html=${coverageFile}
