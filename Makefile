BINDIR = bin
BIN = nightmare-monkey
VERSION ?= undefined
CURRENT_REVISION ?= $(shell git rev-parse --short HEAD)
LDFLAGS = -X 'main.version=$(VERSION)' -X 'main.gitcommit=$(CURRENT_REVISION)'

all: clean gofmt build

build: deps
	go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(BIN) .

GOFMT_RET = .gofmt.txt
gofmt:
	rm -f $(GOFMT_RET)
	gofmt -s -d . | tee $(GOFMT_RET)
	test ! -s $(GOFMT_RET)

deps:
	dep ensure

depup:
	dep ensure -update

clean:
	rm -f $(BINDIR)/$(BIN)

.PHONY: build gofmt deps clean
