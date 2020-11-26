SHELL = /bin/bash

# executables
GO = $(shell which go)
GOPATH = $(shell $(GO) env GOPATH)
GOJUNIT = $(GOPATH)/bin/go-junit-report

# commands
.PHONY: format
format:
	$(GO) mod tidy && go fmt ./...

.PHONY: clean
clean:
	$(GO) clean -x -i -testcache

.PHONY: integration-tests
integration-tests:
	$(GO) test ./e2etests

# needs to call `deps` first since e2etests relies on cmd output
.PHONY: ci-integration-tests
ci-integration-tests: deps install-go-junit-report
	$(GO) test -v ./e2etests 2>&1 | $(GOJUNIT) > report.xml
	$(GO) test -v ./e2etests

.PHONY: unit-tests
unit-tests:
	$(GO) test -short ./...

.PHONY: ci-unit-tests
ci-unit-tests: deps install-go-junit-report
	$(GO) test -v -short ./... 2>&1 | $(GOJUNIT) > report.xml
	$(GO) test -v -short ./...

deps:
	$(GO) get -d ./...

install-go-junit-report:
	$(GO) install github.com/jstemmer/go-junit-report