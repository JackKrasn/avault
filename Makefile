.PHONY: build clean test help default release
BIN_DIR := $(CURDIR)/bin
BIN_NAME ?= avault
DIST_DIRS   := find * -type d -exec

# go option
LDFLAGS    := -w -s

# Rebuild the binary if any of these files change

SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum
.EXPORT_ALL_VARIABLES:

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
GIT_TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
SNAPSHOT_VERSION=$(shell autotag -n -p dev -b ${CURRENT_BRANCH})
VERSION=$(shell autotag -n)

LDFLAGS += -X github.com/JackKrasn/avault/internal/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY}
LDFLAGS += -X github.com/JackKrasn/avault/internal/version.BuildDate=${BUILD_DATE}
LDFLAGS += -X github.com/JackKrasn/avault/internal/version.Version=${SNAPSHOT_VERSION}

default: build

help:
	@echo 'Management commands for avault:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo '    make info            Print info'
	@echo '    make release         Build and publish artifacts'
	@echo '    make snapshot        Build snapshot version'
	@echo '    make test            Run tests'

build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	@echo "building ${BIN_NAME} ${SNAPSHOT_VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '$(LDFLAGS)' -o ${BIN_DIR}/${BIN_NAME} ./cmd/avault

release:
	@echo "release ${BIN_NAME} ${VERSION}"
	autotag
	goreleaser release --clean

snapshot:
	@echo "release ${BIN_NAME} ${SNAPSHOT_VERSION}"
	goreleaser --snapshot --rm-dist

get-deps:
	dep ensure

test:
	go test ./...

clean:
	 @echo "deleting ${BIN_DIR}" ./dist
	 @rm -rf '$(BIN_DIR)' ./dist

.PHONY: info
info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"
