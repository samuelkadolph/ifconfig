BUILD_DATE ?= `date -u +"%Y-%m-%dT%H:%M:%SZ"`
CGO_ENABLED ?= 1
GOBIN ?= go
IMAGE_NAME ?= samuelkadolph/ifconfig:latest
VCS_REF ?= `git rev-parse --short HEAD`

default: clean test build

build: bin/ifconfig

bin/%:
	$(GOBIN) build -o bin/$* .

docker-build:
	docker build --build-arg BUILD_DATE=$(BUILD_DATE) --build-arg VCS_REF=$(VCS_REF) --tag $(IMAGE_NAME) .

clean:
	rm -rf bin/*

test:
	$(GOBIN) test .

.PHONY: build docker-build clean test
