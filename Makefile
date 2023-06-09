EXTENSION ?=
DIST_DIR ?= dist/
EXTRA_DIR ?= extra/
GOOS ?= linux
ARCH ?= $(shell uname -m)
BUILDINFOSDET ?=

SOFT_NAME    := ralgo
SOFT_VERSION := $(or $(shell git describe --tags $(git rev-list --tags --max-count=1)), v0.0)
VERSION_PKG   := $(shell echo $(SOFT_VERSION) | sed 's/^v//g')
ARCH          := x86_64
URL           := https://github.com/Rom1-J/Ralgo
DESCRIPTION   := Encode file in pseudo-random sequence.
BUILDINFOS    :=  ($(shell date +%FT%T%z)$(BUILDINFOSDET))
LDFLAGS       := '-X main.version=$(SOFT_VERSION) -X main.buildinfos=$(BUILDINFOS)'

OUTPUT_SOFT := $(DIST_DIR)ralgo-$(SOFT_VERSION)-$(GOOS)-$(ARCH)$(EXTENSION)

.PHONY: vet
vet:
	go vet main.go

.PHONY: prepare
prepare:
	mkdir -p $(DIST_DIR)

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)

.PHONY: small_sample
small_sample:
	dd if=/dev/urandom of=$(EXTRA_DIR)example/1K bs=1K count=1

.PHONY: sample
sample:
	dd if=/dev/urandom of=$(EXTRA_DIR)example/1M bs=1M count=1

.PHONY: build
build: clean prepare
	go build -ldflags $(LDFLAGS) -o $(OUTPUT_SOFT)