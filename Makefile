CC := go
PWD := $(shell pwd)
OUTPUT_DIR := $(PWD)/dist
BASE_PACKAGE := .
RELEASE := toshoshitsu
VERSION ?= 0.1a
ARCH ?= amd64
PLATFORMS := windows linux darwin
OS = $(word 1, $@)

test:
	$(CC) test $(BASE_PACKAGE)/...

clean:
	@rm -rf $(OUTPUT_DIR)

pre:
	@mkdir -p $(OUTPUT_DIR)

release: clean pre $(PLATFORMS)

windows: POSTFIX = .exe

$(PLATFORMS):
	GOOS=$(OS) GOARCH=$(ARCH) $(CC) build -o $(OUTPUT_DIR)/$(RELEASE)-$(VERSION)-$(OS)-$(ARCH)$(POSTFIX) # $(BASE_PACKAGE)/cmd

all: test clean release

.PHONY: test $(PLATFORMS) release pre clean all
