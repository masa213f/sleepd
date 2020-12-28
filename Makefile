OUTPUT_DIR ?= $(CURDIR)

.PHONY: all
all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/sleepd ./src/...

.PHONY: fmt
fmt:
	goimports -w $$(find . -type f -name '*.go' -print)

.PHONY: lint
lint:
	test -z "$$(goimports -l $$(find . -type f -name '*.go' -print) | tee /dev/stderr)"
	go vet ./...

.PHONY: clean
clean:
	rm -f ./sleepd

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
