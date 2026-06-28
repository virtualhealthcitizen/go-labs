# go-labs — task runner for the Go playground.
#
# These targets wrap the standard Go toolchain so contributors (and CI) have a
# single, consistent entry point. Run `make help` for the list.

GO      ?= go
PKGS    := ./...

.DEFAULT_GOAL := help

.PHONY: help build vet test fmt fmt-check tidy run clean check

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

build: ## Compile every package/example.
	$(GO) build $(PKGS)

vet: ## Run go vet across the module.
	$(GO) vet $(PKGS)

test: ## Run all tests.
	$(GO) test $(PKGS)

fmt: ## Format all Go source in place.
	gofmt -w .

fmt-check: ## Fail if any file is not gofmt-clean (used by CI).
	@unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "These files are not gofmt-clean:"; echo "$$unformatted"; exit 1; \
	fi

tidy: ## Sync go.mod / go.sum with the source.
	$(GO) mod tidy

# Run a single example by path, e.g.:
#   make run DIR=12_concurrency/coroutines_examples
run: ## Run one example: make run DIR=<path-to-example>
	@test -n "$(DIR)" || (echo "usage: make run DIR=<path-to-example>"; exit 1)
	$(GO) run ./$(DIR)

clean: ## Remove build artifacts.
	$(GO) clean
	find . -type f \( -name '*.exe' -o -name '*.test' -o -name '*.out' \) -delete

check: fmt-check vet test ## Run the full gate (fmt-check + vet + test).
