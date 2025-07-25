.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

staticcheck: vet
	staticcheck ./...

revive: staticcheck
	revive ./...

build: revive
	go build