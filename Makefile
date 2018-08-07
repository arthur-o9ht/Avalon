GOPATH:=$(CURDIR)
export GOPATH

all:build

build:
	cd $(CURDIR)/hugoSite && hugo 
	cd $(CURDIR)/src/ && go build -o $(CURDIR)/bin/go-serve
local:
	cd $(CURDIR)/hugoSite && hugo server --baseUrl="http://localhost:1313"

