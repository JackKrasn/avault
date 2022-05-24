.PHONY: build clean test help default
BIN_DIR := $(CURDIR)/bin
BIN_NAME ?= avault
DIST_DIRS   := find * -type d -exec

# go option
LDFLAGS    := -w -s

# Rebuild the binary if any of these files change
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

VERSION := $(shell grep "const Version " internal/version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
GIT_TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)

LDFLAGS += -X github.com/JackKrasn/avault/internal/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY}
LDFLAGS += -X github.com/JackKrasn/avault/internal/version.BuildDate=${BUILD_DATE}

default: test

help:
	@echo 'Management commands for avault:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo '    make info            Print info'


build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '$(LDFLAGS)' -o ${BIN_DIR}/${BIN_NAME} ./cmd/avault

get-deps:
	dep ensure

test:
	go test ./...

clean:
	 @echo "deleting ${BIN_DIR}"
	 @rm -rf '$(BIN_DIR)'

.PHONY: info
info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"
