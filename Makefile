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
	go test -v $(BIN_SRC)/...

.PHONY: lint
lint:
	gometalinter \
		--enable=misspell \
		--enable=gofmt \
		--deadline=120s \
		$(BIN_SRC)/...

.PHONY: build-helper
build-helper: clean
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BIN) $(BIN_ENTRYPOINT)

.PHONY: build-workflow
build-workflow: build-helper
	cp $(RESOURCES_DIR)/* $(BUILD_DIR)
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PKG) *

.PHONY: build
build: build-workflow

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: install-env-deps
install-env-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: install
install:
	go get -u github.com/stretchr/testify/assert

.PHONY: fmt
fmt:
	go fmt $(BIN_SRC)/...
