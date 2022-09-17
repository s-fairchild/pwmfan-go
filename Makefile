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
	build=$$(scripts/go_change_check.sh build/pwmfan-go); \
	if [ $$build == "true" ]; then \
		# tags="-tags placetagshere"; \
		GOARCH=arm go build -o build/pwmfan . ;\
	fi ;\