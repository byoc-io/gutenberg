PROJ=gutenberg
ORG_PATH=github.com/byoc-io
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)

DOCKER_REPO=gutenberg
DOCKER_IMAGE=$(PROJ)

$( shell mkdir -p bin )
$( shell mkdir -p release/bin )
$( shell mkdir -p release/images )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin
# Prefer ./bin instead of system packages for things like protoc, where we want
# to use the version orchestra uses, not whatever a developer has installed.
export PATH=$(GOBIN):$(shell printenv PATH)
export GO15VENDOREXPERIMENT=1

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

build: clean bin/gutenberg

bin/gutenberg: check-go-version
	@echo "Building gutenberg"
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/gutenberg

.PHONY: release-binary
release-binary:
	@echo "Releasing binary files"
	@go build -race -o release/bin/gutenberg -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/gutenberg

.PHONY: revendor
revendor:
	@glide up -v
	# glide-vc --use-lock-file --no-tests --only-code

test:
	@echo "Testing"
	@go test -v -cover -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v -cover $(shell go list ./... | grep -v '/vendor/')

testrace:
	@echo "Testing with race detection"
	@go test -v -cover -i --race $(shell go list ./... | grep -v '/vendor/')
	@go test -v -cover --race $(shell go list ./... | grep -v '/vendor/')

vet:
	@echo "Running go tool vet on packages"
	@go vet $(shell go list ./... | grep -v '/vendor/')

fmt:
	@echo "Running gofmt on package sources"
	@go fmt $(shell go list ./... | grep -v '/vendor/')

lint:
	@echo "lint"
	@for package in $(shell go list ./... | grep -v '/vendor/' | grep -v '/api' | grep -v '/server/internal'); do \
      golint -set_exit_status $$package $$i || exit 1; \
	done

.PHONY: docker-image
docker-image: clean
	@echo "Building $(DOCKER_IMAGE):build image"
	@docker build -t $(DOCKER_IMAGE):build --rm -f Dockerfile-build .

	@echo "Compiling binary files with Docker"
	@docker run --rm -v $(PWD)/release/bin:/go/src/$(REPO_PATH)/release/bin $(DOCKER_IMAGE):build

	@echo "Building $(DOCKER_IMAGE) image"
	@docker build -t $(DOCKER_IMAGE) --rm -f Dockerfile .

.PHONY: check-go-version
check-go-version:
	@echo "Checking Golang version"
	@./scripts/check-go-version

.PHONY: clean
clean:
	@echo "Cleaning Binary Folders"
	@rm -rf bin/*
	@rm -rf release/*

testall: testrace vet fmt # lint

FORCE:

.PHONY: test testrace vet fmt lint testall