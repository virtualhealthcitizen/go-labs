# Repository structure

`go-labs` is organised as a **Go playground**: a single module whose lessons live
under numbered topic directories, and where **every runnable example sits in its
own directory as `package main`**. That one rule is what keeps `go build ./...`,
`go vet ./...`, and `go test ./...` green across the whole repo — two `func main`
declarations can never collide because they never share a directory.

## Layout

```
go-labs/
├── go.mod, go.sum         # one module: "go-labs"
├── Makefile               # build / vet / test / fmt / run helpers
├── docs/                  # topic index + this structure guide
├── util/                  # shared, importable helpers (has unit tests)
├── 01_introduction-to-go/ # prose-only lessons (Markdown)
├── 02_basic_constructs.../ # … topic dirs, each holding runnable examples
│   └── 03_overview_of_data_types/
│       ├── integer_types/integer_types.go   # one example == one dir
│       └── strings/strings.go
├── 12_concurrency/
│   ├── coroutines_examples/coroutines_examples.go
│   ├── channel_patterns/channel_patterns.go
│   └── *.md               # topic notes alongside the runnable examples
├── 19_standard_library/
├── 20_testing/            # examples that ship with *_test.go
└── practice/              # free-form scratch space for your own experiments
```

Topic directories are numbered to follow a learning progression (see
[`docs/topics.md`](topics.md)). The numbering has historical gaps (13–18) that
map to topics on the roadmap but not yet covered — see [`../todo.md`](../todo.md).

## Conventions

- **One runnable example per directory.** A directory that contains a `func main`
  must contain exactly one. To add a variant, give it its own subdirectory.
- **`package main` for runnable snippets; a named package for reusable code.**
  Anything meant to be imported (like `util`) lives in its own package and should
  carry tests.
- **`gofmt`-clean, `go vet`-clean.** CI rejects unformatted code. Run `make fmt`
  before committing.
- **Markdown notes sit beside the code they describe**, named after the topic.

## Running an example

```bash
go run ./09_control_structures/05_for_construct      # by directory
make run DIR=09_control_structures/05_for_construct  # same thing via make
```

There is no need for build tags — the historical `-tags <filename>` instructions
were removed when examples were split one-per-directory.
