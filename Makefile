GOBIN ?= go
CGO_ENABLED ?= 1

default: clean test build

build: bin/ifconfig

bin/%:
	$(GOBIN) build -o bin/$* .

clean:
	rm -rf bin/*

test:
	$(GOBIN) test .

.PHONY: build clean test
