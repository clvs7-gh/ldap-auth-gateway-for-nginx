NAME := auth-gateway
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)

PRJDIR := .
SRCS    := $(shell find ./$(PRJDIR) -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
STFLAGS := -a -tags netgo -installsuffix netgo

.PHONY: x64
x64: $(SRCS)
	cd $(PRJDIR); GO111MODULE=on GOOS=linux GOARCH=amd64 packr build $(STFLAGS) $(LDFLAGS) -o $(CURDIR)/bin/$(NAME)-x64

.PHONY: arm
arm: $(SRCS)
	cd $(PRJDIR); GO111MODULE=on GOOS=linux GOARCH=arm packr build $(STFLAGS) $(LDFLAGS) -o $(CURDIR)/bin/$(NAME)-arm

.PHONY: all
all: $(SRCS)
	cd $(PRJDIR); GO111MODULE=on GOOS=linux GOARCH=amd64 packr build $(STFLAGS) $(LDFLAGS) -o $(CURDIR)/bin/$(NAME)-x64
	cd $(PRJDIR); GO111MODULE=on GOOS=linux GOARCH=arm packr build $(STFLAGS) $(LDFLAGS) -o $(CURDIR)/bin/$(NAME)-arm

.PHONY: dep-install
dep-install:
	go get -u github.com/gobuffalo/packr/packr

.PHONY: clean
clean:
	rm -f bin/$(NAME)-*
	go clean -modcache
