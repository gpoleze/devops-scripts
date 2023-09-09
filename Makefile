SHELL = /usr/bin/env bash

build:
	mkdir "build"
	go build ./cmd/aws-ec2-list-instances/aws-ec2-list-instances.go
	mv aws-ec2-list-instances build

clean:
	go clean
	rm -r build

install:
	go install ./cmd/aws-ec2-list-instances/aws-ec2-list-instances.go