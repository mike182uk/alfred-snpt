BIN=snpt-alfred-workflow
BIN_ENTRYPOINT=./cmd
BIN_SRC=./cmd
BUILD_DIR=./build
RESOURCES_DIR =./resources
WORKFLOW_PKG=Snpt.alfredworkflow

.PHONY: test
test: ## Run the tests
	GO111MODULE=on go test -v $(BIN_SRC)/...

.PHONY: lint
lint: ## Lint the source files
	golangci-lint run $(BIN_SRC)/...

.PHONY: build-helper
build-helper: clean ## Build the helper binary
	GO111MODULE=on \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(BIN) $(BIN_ENTRYPOINT)

.PHONY: build-workflow
build-workflow: build-helper ## Build the Alfred workflow (will also build the helper binary)
	cp $(RESOURCES_DIR)/* $(BUILD_DIR)
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PKG) *

.PHONY: build
build: build-workflow ## Alias for 'build-workflow'

.PHONY: clean
clean: ## Clean the workspace
	rm -rf $(BUILD_DIR)

.PHONY: install-tools
install-tools: ## Install tools required by the project
	if [ -z "$(CI)" ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.0; fi

.PHONY: install
install: install-tools ## Install project dependencies (including any required tools)

.PHONY: fmt
fmt: ## Format the source files
	go fmt $(BIN_SRC)/...

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
