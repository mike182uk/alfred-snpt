BIN_NAME=alfred-snpt
BIN_SRC_DIR=./bin
BUILD_DIR=./build
DEV_DIR=./dev
DEV_CACHE_DIR=$$(pwd)/.cache
DEV_DATA_DIR=$$(pwd)/.data
RESOURCES_DIR =./resources

WORKFLOW_BUNDLE_ID=com.mike182uk.snpt
WORKFLOW_NAME=alfred-snpt
WORKFLOW_VERSION=3.0.0
WORKFLOW_PKG_NAME=snpt.alfredworkflow

.PHONY: env
env: ## Print out env vars required for development
	@echo 'export alfred_workflow_bundleid="$(WORKFLOW_BUNDLE_ID)"'
	@echo 'export alfred_workflow_name="$(WORKFLOW_NAME)"'
	@echo 'export alfred_workflow_version="$(WORKFLOW_VERSION)"'
	@echo 'export alfred_workflow_cache="$(DEV_CACHE_DIR)"'
	@echo 'export alfred_workflow_data="$(DEV_DATA_DIR)"'

.PHONY: lint
lint: ## Lint the source files
	cd $(BIN_SRC_DIR) && golangci-lint run ./...

.PHONY: build-bin
build-bin: clean ## Build the binary
	GOOS=darwin \
	GOARCH=amd64 \
	cd $(BIN_SRC_DIR) && go build -o .$(BUILD_DIR)/$(BIN_NAME) ./

.PHONY: build-workflow
build-workflow: build-bin ## Build the Alfred workflow (will also build the binary)
	cp $(RESOURCES_DIR)/* $(BUILD_DIR)
	sed -i '' 's/__BUNDLE_ID__/$(WORKFLOW_BUNDLE_ID)/g' $(BUILD_DIR)/info.plist
	sed -i '' 's/__VERSION__/$(WORKFLOW_VERSION)/g' $(BUILD_DIR)/info.plist
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PKG_NAME) *

.PHONY: build
build: build-workflow ## Alias for 'build-workflow'

.PHONY: clean
clean: ## Clean the workspace
	rm -rf $(BUILD_DIR) $(DEV_CACHE_DIR) $(DEV_DATA_DIR)

.PHONY: install-tools
install-tools: ## Install tools required by the project
	if [ -z "$(CI)" ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.44.2; fi

.PHONY: install
install: install-tools ## Install project dependencies (including any required tools)
	cd $(BIN_SRC_DIR) && go mod download

.PHONY: fmt
fmt: ## Format the source files
	cd $(BIN_SRC_DIR) && go fmt ./...

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
