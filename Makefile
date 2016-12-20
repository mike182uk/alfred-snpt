BUILD_DIR=./build
WORKFLOW_PACKAGE=Snpt.alfredworkflow
WORKFLOW_HELPER_BIN=snpt-alfred-workflow
WORKFLOW_HELPER_SRC_DIR=./workflow-helper/src/
RESOURCES_DIR=./resources

.PHONY: default
default: build

.PHONY: test
test:
	go test -v $(WORKFLOW_HELPER_SRC_DIR)/...

.PHONY: lint
lint:
	gometalinter \
		--disable=errcheck \
		--enable=gofmt \
		$(WORKFLOW_HELPER_SRC_DIR)

.PHONY: build-helper
build-helper: clean
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(WORKFLOW_HELPER_BIN) $(WORKFLOW_HELPER_SRC_DIR)

.PHONY: build-workflow
build-workflow: build-helper
	cp resources/* $(BUILD_DIR)
	cd $(BUILD_DIR) && zip -rq $(WORKFLOW_PACKAGE) *

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
	go fmt $(WORKFLOW_HELPER_SRC_DIR)/...
