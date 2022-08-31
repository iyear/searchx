GOOS = $(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
PKG = github.com/iyear/searchx

.PHONY: build
build:
	goreleaser build --rm-dist --single-target --snapshot
	@echo "go to '.searchx/dist' directory to see the package!"

.PHONY: packaging
packaging:
	goreleaser release --skip-publish --auto-snapshot --rm-dist
	@echo "go to '.searchx/dist' directory to see the packages!"

