BUILD_DATE ?= `date -u +"%Y-%m-%dT%H:%M:%SZ"`
CGO_ENABLED ?= 1
DOCKER_IMAGE ?= samuelkadolph/openvpn-monitor
GOBIN ?= go
VCS_REF ?= `git rev-parse --short HEAD`

default: clean test build

docker-build:
	docker build --build-arg $(BUILD_DATE) --build-arg $(VCS_REF) --tag $(DOCKER_IMAGE) .

build: bin/ifconfig

bin/%:
	$(GOBIN) build -o bin/$* .

clean:
	rm -rf bin/*

test:
	$(GOBIN) test .

.PHONY: build docker-build clean test
