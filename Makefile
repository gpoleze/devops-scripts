SHELL = /usr/bin/env bash

build:
	mkdir "build"
	find cmd -name "*.go" -exec go build {} \;
	ls cmd | xargs -n1 -I% mv % build
clean:
	go clean
	rm -r build

install:
	find cmd -name "*.go" -exec go install {} \;