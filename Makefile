BIN=alfred-snpt
BIN_SRC=./bin
BUILD_DIR=./build
RESOURCES_DIR =./resources
WORKFLOW_PKG=snpt.alfredworkflow

.PHONY: test
test: ## Run the tests
	cd $(BIN_SRC) && go test -v ./...

.PHONY: lint
lint: ## Lint the source files
	cd $(BIN_SRC) && golangci-lint run ./...

.PHONY: build-bin
build-bin: clean ## Build the binary
	GOOS=darwin \
	GOARCH=amd64 \
	cd $(BIN_SRC) && go build -o .$(BUILD_DIR)/$(BIN) ./

.PHONY: build-workflow
build-workflow: build-bin ## Build the Alfred workflow (will also build the binary)
	cp $(RESOURCES_DIR)/* $(BUILD_DIR)
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PKG) *

.PHONY: build
build: build-workflow ## Alias for 'build-workflow'

.PHONY: clean
clean: ## Clean the workspace
	rm -rf $(BUILD_DIR)

.PHONY: install-tools
install-tools: ## Install tools required by the project
	if [ -z "$(CI)" ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.44.2; fi

.PHONY: install
install: install-tools ## Install project dependencies (including any required tools)
	cd $(BIN_SRC) && go mod download

.PHONY: fmt
fmt: ## Format the source files
	cd $(BIN_SRC) && go fmt ./...

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
