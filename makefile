all: test build run clean

VERSION := $(shell git describe --tags)
GOLDFLAGS += -X main.Version=$(VERSION)
GOFLAGS = -ldflags "$(GOLDFLAGS)"
run:
	@echo "--------RUNNING------------------"
	@./bin/tmg && echo

dev:
	@go run . --config=./test/demo.tmg && echo

watch:
	@air

build:
	@echo "--------BUILDING-----------------"
	@go build -v -o bin/tmg $(GOFLAGS) ./app.go

clean:
	@echo "--------CLEANING ARTIFACT--------"
	@rm -rf bin/

test:
	@echo "--------TEST STAGE---------------"
	@go test ./...

.PHONY: build test clean all dev watch
