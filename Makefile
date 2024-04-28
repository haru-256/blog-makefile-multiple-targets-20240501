ifndef GOARCH
	GOARCH=$(shell go env GOARCH)
endif

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif

COMMAND_DIRS := $(wildcard cmd/*)
BUILD_TARGETS := $(addprefix build-,$(notdir $(COMMAND_DIRS)))

.PHONY: $(BUILD_TARGETS)
$(BUILD_TARGETS): build-%:
	CGO_ENABLED=0 GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/$* -mod=vendor cmd/$*/main.go

.PHONY: lint
lint: ## Run Linter
	golangci-lint run --config=./.golangci.yaml ./...

.PHONY: fmt
fmt: ## Run formatter
	go fmt ./...
