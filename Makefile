GOPATH = $(shell echo $$GOPATH)
BUILD_DIR = _output

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: deps
deps:
	go get ./...

.PHONY: build
build: $(BUILD_DIR)/rknresolver

test: test-deps
	$(GOPATH)/bin/ginkgo -r --randomizeAllSpecs -p -nodes=4

test-deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

$(BUILD_DIR)/rknresolver: $(BUILD_DIR) deps
	go build -o $@ ./cmd/rknresolver