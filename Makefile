GO		        := $(shell which go)
GOPATH        := $(shell go env GOPATH)
GOBIN         := $(GOPATH)/bin
GOLINT        := $(GOBIN)/golint
COVERAGE_FILE	:= coverage.out
MOCKERY       := $(shell which mockery)

# main tasks
.PHONY: run
run:
	$(GO) run cmd/main.go

.PHONY: test
test: clean lint vet unit

.PHONY: coverage-html
coverage-html: unit
	$(GO) tool cover -html=$(COVERAGE_FILE)

# sub tasks
.PHONY: clean
clean:
	rm -rf $(COVERAGE_FILE)

.PHONY: lint-install
lint-install:
	test -e $(GOLINT) || $(GO) get -u golang.org/x/lint/golint

.PHONY: lint
lint: lint-install
	$(GO) list ./... | grep -v /vendor/ | xargs -L1 $(GOLINT) -set_exit_status

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: unit
unit:
	$(GO) test -race -v ./... -coverprofile=$(COVERAGE_FILE)

.PHONY: coverage
coverage: unit
	$(GO) tool cover -func=$(COVERAGE_FILE)

.PHONY: mockery-install
mockery-install:
	test -e $(MOCKERY) || $(GO) get github.com/vektra/mockery/v2/.../

.PHONY: generate
generate: mockery-install
	$(GO) generate ./...

build:
	go build -o bin/coinmarket-server cmd/main.go && ./bin/coinmarket-server

build-lambda:
	CGO_ENABLED=0 GOOS=linux go build -o bin/coinmarket-lambda cmd/main.go

scan:
	gosec -exclude-generated -nosec=false ./... ./...