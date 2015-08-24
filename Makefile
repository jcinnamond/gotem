RELEASE=releases
PROJECT=github.com/jcinnamond/gotem
ENV=/usr/bin/env
BUILD=go build
SRC=gotem.go version.go
NAME=$(shell go run $(SRC) -v)

.PHONY: release

release: $(SRC) | $(RELEASE)/
	cd $(RELEASE) && $(ENV) GOOS=linux GOARCH=amd64 $(BUILD) -o "$(NAME)-linux_amd64" $(PROJECT)
	cd $(RELEASE) && $(ENV) GOOS=darwin GOARCH=amd64 $(BUILD) -o "$(NAME)-osx" $(PROJECT)

$(RELEASE)/:
	mkdir -p $@
