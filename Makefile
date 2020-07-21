.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/slackme slackme.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/slackme-lin slackme.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/slackme.exe slackme.go

.DEFAULT_GOAL := build
