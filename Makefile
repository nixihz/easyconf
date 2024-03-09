.PHONY: build
# generate build
build:
	go build -o bin/easyconf main.go

.PHONY: install
# install test
install:
	go install ./
