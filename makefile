PACKAGES := ./ ./level ./message ./send

build: 
	go build -v

lint:
	gofmt -l $(PACKAGES)
	go vet $(PACKAGES)

test:build lint
	go test -v -check.v $(PACKAGES)
