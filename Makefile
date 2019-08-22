SHELL := $(shell which bash)
ENV := /usr/bin/env
PROJECT_NAME := "go-api-benchmark"
PKG := "github.com/juliendoutre/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
OSARCH := "linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386"

.PHONY: all dep build clean test

all: clean dep vet fmt test build

clean:
	@rm -f c.out
	@rm -f coverage.html
	@rm -rf bin/
	@rm -rf vendor/

dep:
	go get -v -u github.com/golang/dep/cmd/dep && \
	go get -v -u github.com/mitchellh/gox && \
	go get -v -u github.com/mattn/goveralls && \
	dep ensure -v -vendor-only

fmt:
	go fmt ${PKG_LIST}

vet:
	go vet ${PKG_LIST}

test: dep
	go test -cover -coverprofile=c.out ${PKG_LIST}
	go tool cover -html=c.out -o .coverage.html
	@rm -f c.out

build: dep
	env GOOS=${GOOS} GOARCH=${GOARCH} go build -i -o "bin/$(PROJECT_NAME)" -v $(PKG)

cross: dep
	gox -osarch=$(OSARCH) -output "bin/$(PROJECT_NAME)_{{.OS}}_{{.Arch}}"
