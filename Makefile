PROJECT=github.com/smartdigits/gohealth
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) GO15VENDOREXPERIMENT=1 $(GO)

.PHONY: test all clean dependencies setup example

all: setup test

clean:
	rm -fr src/
	rm -fr bin/
	rm -fr pkg/

setup:
	mkdir -p src/$(PROJECT)
	rm -fr src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)

dependencies:
	$(GOCMD) get $(PROJECT)

test:
	$(GOCMD) test $(PROJECT)
