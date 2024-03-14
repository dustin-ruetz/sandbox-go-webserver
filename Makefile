# If no target is specified (i.e. the `make` command is run by itself), default to running `make dev`
# so that the codebase is cleaned, validated and built prior to starting the webserver.
.DEFAULT_GOAL := dev

build: clean validate
	go build ./main.go
.PHONY: build

clean:
	rm -f ./main
.PHONY: clean

dev: build
	./main
.PHONY: dev

fmt:
	go fmt ./...
.PHONY: fmt

lint:
	golangci-lint run ./...
.PHONY: lint

# Run `make lintall` to have golangci-lint enable _all_ linters when checking the codebase.
# This target is deliberately excluded from `validate` because the output is very noisy.
lintall:
	golangci-lint run --enable-all ./...
.PHONY: lintall

test:
	go test ./...
.PHONY: test

validate: fmt lint test vet
.PHONY: validate

vet:
	go vet ./...
.PHONY: vet
