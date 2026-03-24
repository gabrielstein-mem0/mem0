VERSION ?= dev
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -s -w \
	-X github.com/mem0ai/mem0/mem0-cli/cmd.version=$(VERSION) \
	-X github.com/mem0ai/mem0/mem0-cli/cmd.commit=$(COMMIT) \
	-X github.com/mem0ai/mem0/mem0-cli/cmd.date=$(DATE) \
	-X github.com/mem0ai/mem0/mem0-cli/internal/api.version=$(VERSION)

.PHONY: build clean fmt lint vet test

build:
	go build -ldflags "$(LDFLAGS)" -o mem0 .

clean:
	rm -f mem0

fmt:
	gofmt -w .

lint: fmt vet
	golangci-lint run ./...

vet:
	go vet ./...

test:
	go test ./...
