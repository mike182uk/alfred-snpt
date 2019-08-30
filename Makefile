BIN=snpt-alfred-workflow
BIN_ENTRYPOINT=./cmd
BIN_SRC=./cmd
BUILD_DIR=./build
RESOURCES_DIR =./resources
WORKFLOW_PKG=Snpt.alfredworkflow

.PHONY: default
default: build

.PHONY: test
test:
	GO111MODULE=on go test -v $(BIN_SRC)/...

.PHONY: lint
lint:
	GO111MODULE=on golangci-lint run $(BIN_SRC)/...

.PHONY: build-helper
build-helper: clean
	GO111MODULE=on \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(BIN) $(BIN_ENTRYPOINT)

.PHONY: build-workflow
build-workflow: build-helper
	cp $(RESOURCES_DIR)/* $(BUILD_DIR)
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PKG) *

.PHONY: build
build: build-workflow

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: install-tools
install-tools:
	GO111MODULE=on \
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: install
install: install-tools

.PHONY: fmt
fmt:
	go fmt $(BIN_SRC)/...
