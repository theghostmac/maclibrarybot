# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

# Project parameters
BINARY_NAME=maclibrarybot
BINARY_UNIX=$(BINARY_NAME)_unix

# Default port for commands
PORT=8080  # Replace with the actual port your server listens on.

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GORUN) ./cmd/main.go

deps:
	$(GOCMD) mod download

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/main.go

# New commands
start: build
	./$(BINARY_NAME) start &

show-port:
	sudo lsof -i TCP:$(PORT)

get-pid:
	pgrep $(BINARY_NAME)

stop:
	sudo kill -9 $$(pgrep $(BINARY_NAME))

.PHONY: all test clean run deps build-linux start show-port get-pid stop%