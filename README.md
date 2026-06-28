# go-labs

A hands-on **Go playground** — a single module full of small, runnable examples
that walk through the language from first principles (variables, control flow)
through to concurrency, the standard library, and testing. Every snippet is a
self-contained program you can run, read, and tinker with.

- **Go**: 1.23+ (the module targets `go 1.23`; see [`go.mod`](go.mod))
- **Editor**: any — the repo is editor-agnostic (no IDE files are tracked)

## Quick start

```bash
git clone https://github.com/virtualhealthcitizen/go-labs
cd go-labs

go build ./...        # compile every example
go test ./...         # run the tests
go run ./09_control_structures/05_for_construct   # run one example
```

If you have `make`:

```bash
make            # list available targets
make check      # full gate: gofmt-check + vet + test
make run DIR=12_concurrency/coroutines_examples
```

## How it's organised

Lessons live under numbered topic directories that follow a learning
progression, and **every runnable example sits in its own directory** as
`package main`. That single rule keeps `go build ./...` green — two `func main`
declarations can never collide. Reusable, importable code (e.g. [`util/`](util))
lives in its own named package and ships with tests.

See [`docs/STRUCTURE.md`](docs/STRUCTURE.md) for the full layout and conventions,
and [`docs/topics.md`](docs/topics.md) for the topic roadmap.

> **Note:** earlier revisions documented a `go run -tags <filename>` workflow for
> picking one example out of a shared directory. That is no longer needed (or
> supported) — examples were split one-per-directory, so you select an example by
> its path, not a build tag.

## Working with Go modules

This repo is one module named `go-labs`. The commands below are the ones you'll
reach for most; the official [modules reference](https://go.dev/ref/mod) has the
rest.

| Command          | What it does                                            |
| ---------------- | ------------------------------------------------------- |
| `go mod tidy`    | Add missing / drop unused dependencies in `go.mod`      |
| `go build ./...` | Compile all packages                                    |
| `go vet ./...`   | Report likely mistakes the compiler doesn't catch       |
| `go test ./...`  | Run all tests                                            |
| `gofmt -w .`     | Format all source in place                              |
| `go run ./<dir>` | Build and run a single example without leaving a binary |

## Contributing examples

1. Create a new directory under the relevant topic (one example per directory).
2. Write `package main` with a single `func main`; keep it focused on one idea.
3. Run `make check` (or `gofmt -w . && go vet ./... && go test ./...`) — CI runs
   the same gate on every push and pull request.
4. Add a short Markdown note beside the code if the concept needs prose.

See [`todo.md`](todo.md) for the current backlog, including topics still to be
covered and ecosystem deep-dives.

## Additional resources

- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Standard library docs](https://pkg.go.dev/std)
