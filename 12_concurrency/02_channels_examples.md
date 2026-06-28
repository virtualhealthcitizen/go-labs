# Channels

Channels are Go's primary mechanism for communication between goroutines:
*"Don't communicate by sharing memory; share memory by communicating."*

The runnable example in [`channels/channels.go`](channels/channels.go) walks
the core patterns in order:

1. **Unbuffered channels** — a send blocks until another goroutine receives.
2. **Buffered channels** — sends don't block until the buffer is full (`len`/`cap`).
3. **Channel directions** — `chan<-` (send-only) and `<-chan` (receive-only)
   enforced at compile time.
4. **Closing & ranging** — `close(ch)` signals "no more values"; `range` stops
   once the channel is closed and drained.
5. **The comma-ok receive** — `v, ok := <-ch` distinguishes a real send
   (`ok == true`) from the zero value of a closed channel (`ok == false`); this
   is what `range` checks under the hood.
6. **Worker pool** — fan out jobs to N workers, fan in their results.
7. **`select`** — with `time.After` timeouts and a non-blocking `default`.

Run it:

```shell
go run ./12_concurrency/channels
# or, from the repo root:
make run DIR=12_concurrency/channels
```

For higher-level composition (generic fan-in, pipelines, context cancellation),
see [`channel_patterns/`](channel_patterns/channel_patterns.go).
