# go-labs — Backlog & Burndown

Autonomous burndown backlog for the Go playground. **One validated item per burn,
shipped via PR**, with every gate green:

```
gofmt -l .   # must be empty
go vet ./...
go build ./...
go test ./...
```

(or simply `make check`). Mark items `[x]` done / `[~]` partial with a one-line
note, and prepend a dated entry to the Burndown Log. Prefer items that make the
repo a better learning environment: more runnable examples, clearer notes, and
hands-on tours of the wider Go ecosystem.

## Priority for the next rounds

### High
- [x] **Flesh out the channels example** — added runnable
  `12_concurrency/channels/channels.go` (unbuffered vs buffered, directionality,
  `close` + `range`, comma-ok receive, worker pool, `select` with timeout +
  `default`); rewrote `02_channels_examples.md` to point at the runnable file
  instead of embedding a stale copy. Verified: `go run ./12_concurrency/channels`
  prints all 7 sections. PR #2.
- [ ] **Per-topic READMEs** — several topic dirs (02, 03, 05, 09…) have no
  `README.md`. Add a short one per topic: what it covers, which examples to run,
  links to the relevant notes. Use `06_maps`/`07_structs_and_methods` as the
  template (they already have one).  ← next
- [ ] **Error-handling topic (new `13_error_handling/`)** — fill the first
  numbering gap. Examples: sentinel errors, `errors.Is`/`errors.As`, wrapping with
  `%w`, custom error types, `defer`/`panic`/`recover`. Cross-link from
  `docs/topics.md`.

### Medium
- [ ] **Interfaces & type system topic (new `14_interfaces/`)** — interface
  satisfaction, embedding/composition, type assertions, type switches, the empty
  interface and `any`, and a note on accept-interfaces/return-structs.
- [ ] **More unit tests** — only `util/` and `20_testing/basics/example` have
  tests. Add table-driven tests to a few pure-function examples (e.g.
  `04_functions/01_basic_syntax/average`) so the repo demonstrates testing
  alongside each concept.
- [ ] **`go run` index / runner script** — a small `scripts/list-examples.sh` (and
  `.bat`) that prints every runnable example dir, so newcomers can discover what's
  available without spelunking. Optionally wire `make list`.
- [ ] **Replace `build_all.{sh,bat}`** — the `find … -execdir go build` scripts are
  superseded by `go build ./...` / `make build`. Either delete them or convert to
  thin wrappers that call the Makefile, and update any references.

### Later
- [ ] **Module path → repo URL** — consider renaming the module from `go-labs` to
  `github.com/virtualhealthcitizen/go-labs` so examples can import shared packages
  by their canonical path (matches the README's published-module guidance). Update
  imports of `util` if any are added.
- [ ] **File I/O & system ops topic (new `15_file_io/`)** — `os`, `io`, `bufio`,
  `io.Reader`/`io.Writer`, `path/filepath`, temp files.
- [ ] **Generics topic (new `16_generics/`)** — type parameters, constraints
  (`comparable`, `constraints`), generic data structures. (`channel_patterns.go`
  already previews generics via `fanIn[T any]`.)
- [ ] **`golangci-lint` in CI** — add a lint job alongside the build/vet/test
  workflow once the example set is stable.

## Ecosystem discovery cycles

Each item is a **discovery cycle**: survey the option(s), then land a tiny,
self-contained, runnable example (its own module-internal dir, or a sub-module if
heavy deps are involved) plus a short note on when to reach for it. Keep
dependencies isolated so the core learning examples stay dependency-light.

- [ ] **HTTP & web** — stdlib `net/http` first (server, `ServeMux`, client), then a
  comparison cycle across **chi**, **Gin**, **Echo**, and **Fiber** (routing,
  middleware, JSON, when to pick which).
- [ ] **CLI tooling** — `flag` (stdlib) → **cobra** (+ **viper** for config) and
  **urfave/cli**; build the same small command three ways.
- [ ] **Structured logging** — stdlib **`log/slog`** (Go 1.21+) as the baseline,
  contrasted with **zerolog** and **zap** (handlers, levels, structured fields).
- [ ] **Testing & mocks** — **testify** (assert/require/suite), **gomock** /
  **mockery**, golden-file testing, and fuzzing (`go test -fuzz`).
- [ ] **Database access** — `database/sql` + a driver, then **sqlc** (codegen),
  **sqlx**, and **GORM**; a note on the tradeoffs (raw SQL vs ORM).
- [ ] **Concurrency helpers** — `golang.org/x/sync` (**errgroup**, `semaphore`,
  `singleflight`) building on the existing `12_concurrency` examples.
- [ ] **Config & env** — **viper**, **envconfig**, and `os`/`flag` patterns for
  12-factor-style configuration.
- [ ] **HTTP clients & resilience** — `net/http` client tuning, **resty**, retries,
  timeouts, and context propagation.
- [ ] **Serialization** — stdlib `encoding/json`, plus **protobuf**/**gRPC** and a
  note on `encoding/gob` and YAML.
- [ ] **Build & dev workflow** — **air** (live reload), **golangci-lint**,
  **goreleaser**, and `go work` (workspaces) for multi-module setups.

## Burndown Log

- 2026-06-28 — Sprint 1: Flesh out the channels example. Extracted the channels
  walkthrough that previously lived only as a fenced code block inside
  `02_channels_examples.md` into a runnable `12_concurrency/channels/channels.go`,
  and added an explicit comma-ok-receive section (`v, ok := <-ch` on a closed
  channel) that the prose listed but the code omitted. The 7 sections cover
  unbuffered/buffered channels, directional params, close + range, comma-ok,
  a 3-worker pool, and `select` with `time.After` + non-blocking `default`.
  Rewrote the `.md` into a concise topic note linking the runnable file (per the
  one-example-per-dir convention in `docs/STRUCTURE.md`). Gate: `gofmt -l` empty,
  `go vet ./...`, `go build ./...`, `go test ./...` green; `go run` prints all
  sections correctly. PR #2.

- 2026-06-28 — Sprint 0: Consolidation into a clean Go playground. Split the two
  directories that held multiple `func main` declarations into one-example-per-dir
  (`09_control_structures/03_switch_case_construct/{switch_case_construct_1,
  switch_case_construct_2}`, `12_concurrency/{coroutines_examples,channel_patterns}`)
  so `go build ./...` passes; removed an empty `02_channels_examples.go` stub.
  Replaced the dead `func main` in `package util` with real, importable helpers
  (`Banner`/`Section`) + table-driven tests. Bumped `go.mod` 1.19 → 1.23 and ran
  `go mod tidy` (the 1.22+ loop-var change also cleared the prior `go vet`
  loop-capture warnings). Ran `gofmt -w .` across the whole module (it had never
  been formatted). Added a `Makefile` (build/vet/test/fmt/run/check), GitHub
  Actions CI (gofmt + vet + build + test), `docs/STRUCTURE.md`, and rewrote the
  README (removed the obsolete `-tags` workflow). Seeded this backlog.
  Gate: `gofmt -l .` empty, `go vet ./...`, `go build ./...`, `go test ./...` all
  green.
