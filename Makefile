VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.version=$(VERSION)
BIN     := bin/qbo

.PHONY: build test lint clean fmt vet

build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/qbo

test:
	go test -race ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

clean:
	rm -rf bin/

install: build
	cp $(BIN) $(GOPATH)/bin/qbo 2>/dev/null || cp $(BIN) ~/go/bin/qbo
