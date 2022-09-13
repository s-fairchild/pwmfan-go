ONESHELL:
SHELL = /bin/bash

gochecks:
	go mod tidy
	go fmt
	go vet

build: gochecks
	if [[ ! -d build ]]; then \
		mkdir build ;\
	fi ;\
	cd build ;\
	GOARCH=arm go build ../