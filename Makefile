GOPATH:=$(CURDIR)
export GOPATH

all:build

build:
	cd $(CURDIR)/src/ && go build -o $(CURDIR)/bin/go-serve
