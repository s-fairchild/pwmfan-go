ONESHELL:
SHELL = /bin/bash

gochecks:
	go mod tidy
	go fmt
	go vet

gotests: gochecks
	go test ./...

build: gotests
	if [[ ! -d build ]]; then \
		mkdir build ;\
	fi ;\
	build=$$(scripts/go_change_check.sh build/pwmfan-go); \
	if [ $$build == "true" ]; then \
		# tags="-tags placetagshere"; \
		GOARCH=arm go build -o build/pwmfan . ;\
	fi ;\

source: gotests
	currentBranch=$$(git branch | grep '*' | tr -d '[:space:]' | tr -d '*') ;\
	highestTagVer=$$(git tag -l | sort -V) ;\
	git checkout $$highestTagVer ;\
	if [[ ! -d build/source ]]; then \
		mkdir -p build/source ;\
	fi ;\
	sourceArchive="build/source/pwmfan-$$highestTagVer.tar.gz" ;\
	tar czvf $$sourceArchive --exclude=.git* --exclude=build ./ ;\
	git checkout $$currentBranch ;\

gh-create-release:
	# TODO script release creation

gh-upload:
	latestSource=$$(ls -tr build/source/) ;\
	latestSourceTag=$$(ls -tr build/source/ | cut -d '-' -f 2 | sed 's/.tar.gz//') ;\
	gh release upload $$latestSourceTag build/source/$$latestSource ;\