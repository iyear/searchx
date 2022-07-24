GOOS = $(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
VERSION = $(shell git describe --tags --always --match='v*' | cut -c 1-)
PKG = github.com/iyear/searchx
GO_LDFLAGS = -X '$(PKG)/global.Version=$(VERSION)'
BUILD_DIR := ./.searchx

.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@echo "Version: $(VERSION)"
	go build -ldflags "$(GO_LDFLAGS)" -o "$(BUILD_DIR)/searchx-${GOOS}-${GOARCH}"
	@echo "'searchx' has been built in the '.searchx' directory)!"

.PHONY: packaging
packaging:
	goreleaser release --skip-publish --snapshot --rm-dist
	@echo "go to '.searchx/dist' directory to see the packages!"

