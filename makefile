# project configuration
name := grip
buildDir := build
packages := logging grip # TODO: add other packages when there are tests
projectPath := github.com/tychoish/$(name)


# declaration of dependencies
lintDeps := github.com/alecthomas/gometalinter
testDeps := github.com/stretchr/testify
deps := github.com/coreos/go-systemd/journal


# linting configuration
levelsRegex :=	(Catch.*|Default.*|Emergency.*|Alert.*|Critical.*|Error.*|Warning.*|Notice.*|Debug.*|Info.*)
lintExclusion := --exclude="exported method Grip\.$(levelsRegex)"
lintExclusion += --exclude="exported function $(levelsRegex)"
lintExclusion += --exclude="exported method InternalSender\..*"
lintExclusion += --exclude="package comment should be of the form \"Package grip \.\.\.\""
# end linting configuration


# implementation details for being able to lazily
gopath := $(shell go env GOPATH)
deps := $(addprefix $(gopath)/src/,${deps})
lintDeps := $(addprefix $(gopath)/src/,${lintDeps})
testDeps := $(addprefix $(gopath)/src/,${testDeps})
$(gopath)/src/%:
	@-[ ! -d $(gopath) ] && mkdir -p $(gopath) || true
	go get $(subst $(gopath)/src/,,$@)
# end dependency installation tools


# userfacing targets for basic build/test/lint operations
phony := lint build test coverage coverage-html
lint:
	$(gopath)/bin/gometalinter --deadline=20s --disable=gotype $(lintExclusion) ./...
build:
	go build ./...
test:$(foreach target,$(packages),$(buildDir)/test.$(target).out)
coverage:$(foreach target,$(packages),$(buildDir)/coverage.$(target).out)
coverage-html:$(foreach target,$(packages),$(buildDir)/coverage.$(target).html)
# end front-ends


# implementation details for building the binary and creating a
# convienent link in the working directory
$(gopath)/src/$(projectPath):
	rm -f $@
	mkdir -p `dirname $@`
	ln -s $(shell pwd) $@
$(name):$(buildDir)/$(name)
	[ -L $@ ] || ln -s $< $@
$(buildDir)/$(name):$(gopath)/src/$(projectPath)
	go build -o $@ main/$(name).go
phony += $(buildDir)/$(name)
# end main build


# convenience targets for runing tests and coverage tasks on a
# specific package.
test-%:
	$(MAKE) $(buildDir)/test.$*.out
coverage-%:
	$(MAKE) $(buildDir)/coverage.$*.out
coverage-html-%:
	$(MAKE) $(buildDir)/coverage.$*.html
phony += $(foreach target,$(packages),test-$(target))
phony += $(foreach target,$(packages),coverage-$(target))
phony += $(foreach target,$(packages),coverage-html-$(target))
# end convienence targets


# implementation for package coverage and test running
$(buildDir)/coverage.%.html:$(buildDir)/coverage.%.out
	go tool cover -html=$< -o $@
$(buildDir)/coverage.%.out:test-deps
	go test -covermode=count -coverprofile=$@ $(projectPath)/$*
	@-[ -f $@ ] && go tool cover -func=$@ | sed 's%$(projectPath)/%%' | column -t
$(buildDir)/coverage.$(name).out:test-deps
	go test -covermode=count -coverprofile=$@ $(projectPath)
	@-[ -f $@ ] && go tool cover -func=$@ | sed 's%$(projectPath)/%%' | column -t
$(buildDir)/test.%.out:test-deps
	go test -v ./$* >| $@; exitCode=$$?; cat $@; [ $$exitCode -eq 0 ]
$(buildDir)/test.$(name).out:test-deps
	go test -v ./ >| $@; exitCode=$$?; cat $@; [ $$exitCode -eq 0 ]
# end coverage rports


# targets to install dependencies
deps:$(deps)
test-deps:$(testDeps)
lint-deps:$(lintDeps)
	$(gopath)/bin/gometalinter --install
clean:
	rm -rf $(deps) $(lintDeps) $(testDeps)
.PHONY:$(phony)
