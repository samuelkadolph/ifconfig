BUILD_DATE ?= `date -u +"%Y-%m-%dT%H:%M:%SZ"`
GO ?= go
IMAGE_NAME ?= samuelkadolph/ifconfig:latest
VCS_REF ?= `git rev-parse --short HEAD`

default: build

build: bin/ifconfig

bin/%: $(filter-out %test.go, $(wildcard *.go))
	$(GO) build -o bin/$* .

docker-build:
	docker build --build-arg BUILD_DATE=$(BUILD_DATE) --build-arg VCS_REF=$(VCS_REF) --tag $(IMAGE_NAME) .

docker-publish:
	docker push $(IMAGE_NAME)

clean:
	rm -rf bin/*

test:
	$(GO) test .

.PHONY: build docker-build docker-publish clean test
