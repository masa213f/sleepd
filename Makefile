OUTPUT_DIR ?= $(CURDIR)

VERSION ?= devel
IMAGE_TAG ?= $(VERSION)

.PHONY: all
all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/sleepd -trimpath -ldflags "-s -w -X main.Version=$(VERSION)" ./src/...

.PHONY: fmt
fmt:
	goimports -w $$(find . -type f -name '*.go' -print)

.PHONY: lint
lint:
	test -z "$$(goimports -l $$(find . -type f -name '*.go' -print) | tee /dev/stderr)"
	staticcheck ./...
	go vet ./...

.PHONY: test
test: lint

.PHONY: image
image:
	$(MAKE) build VERSION=$(VERSION) OUTPUT_DIR=./image
	cp LICENSE ./image
	docker build -t sleepd:devel image

.PHONY: tag
tag:
	docker tag sleepd:devel $(IMAGE_PREFIX)sleepd:$(IMAGE_TAG)

.PHONY: push
push:
	docker push $(IMAGE_PREFIX)sleepd:$(IMAGE_TAG)

.PHONY: clean
clean:
	rm -f ./sleepd
	rm -rf ./image/sleepd ./image/LICENSE

.PHONY: setup
setup: goimports staticcheck

.PHONY: goimports
goimports: 
	if ! which goimports >/dev/null; then \
		cd /tmp; env GOFLAGS= GO111MODULE=on go get golang.org/x/tools/cmd/goimports; \
	fi

staticcheck:
	if ! which staticcheck >/dev/null; then \
		cd /tmp; env GOFLAGS= GO111MODULE=on go get honnef.co/go/tools/cmd/staticcheck; \
	fi
