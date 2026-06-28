// channel_patterns.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//
// 1. FAN-IN: merge multiple input channels into a single channel
//

// fanIn merges multiple input channels into one output channel.
func fanIn[T any](inputs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	wg.Add(len(inputs))
	for _, ch := range inputs {
		ch := ch // capture
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	// Close out after all inputs are drained.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

//
// 2. FAN-OUT (worker pool): multiple workers reading from one channel
//

func worker2(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		// Simulate work
		time.Sleep(150 * time.Millisecond)
		results <- fmt.Sprintf("worker %d processed job %d", id, j)
	}
}

//
// 3. PIPELINE: chain multiple stages using channels
//

// gen is a generator stage that emits the given numbers.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square reads ints, squares them, and sends them on.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// double reads ints, doubles them, and sends them on.
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

//
// 4. DONE / QUIT CHANNEL: cooperative cancellation without context
//

// tickerWithDone sends "tick" periodically until done is closed.
func tickerWithDone(done <-chan struct{}, interval time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				out <- fmt.Sprintf("tick at %v", t.Format("15:04:05.000"))
			}
		}
	}()
	return out
}

//
// 5. CONTEXT CANCELLATION: idiomatic cancellation in Go
//

func doWorkWithContext(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stopping:", ctx.Err())
			return
		default:
			fmt.Println("worker", id, "doing work")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	// ------------------------------------------------------------------
	// 1. FAN-IN DEMO
	// ------------------------------------------------------------------
	fmt.Println("=== 1. Fan-in (merge multiple channels) ===")

	a := make(chan string)
	b := make(chan string)

	// Producer A
	go func() {
		defer close(a)
		for i := 1; i <= 3; i++ {
			a <- fmt.Sprintf("A-%d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Producer B
	go func() {
		defer close(b)
		for i := 1; i <= 3; i++ {
			b <- fmt.Sprintf("B-%d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	merged := fanIn(a, b)
	for msg := range merged {
		fmt.Println("merged:", msg)
	}
	fmt.Println()

	// ------------------------------------------------------------------
	// 2. FAN-OUT / WORKER POOL DEMO
	// ------------------------------------------------------------------
	fmt.Println("=== 2. Fan-out (worker pool) ===")

	jobs := make(chan int)
	results := make(chan string)

	// Start workers (fan-out)
	const numWorkers = 3
	for i := 1; i <= numWorkers; i++ {
		go worker2(i, jobs, results)
	}

	// Send jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Collect results
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()

	// ------------------------------------------------------------------
	// 3. PIPELINE DEMO
	// ------------------------------------------------------------------
	fmt.Println("=== 3. Pipeline (gen -> square -> double) ===")

	// Build the pipeline: numbers -> square -> double
	nums := gen(1, 2, 3, 4, 5)
	sq := square(nums)
	dbl := double(sq)

	for v := range dbl {
		fmt.Println("output:", v)
	}
	fmt.Println()

	// ------------------------------------------------------------------
	// 4. DONE / QUIT CHANNEL DEMO
	// ------------------------------------------------------------------
	fmt.Println("=== 4. Done / quit channel ===")

	done := make(chan struct{})
	ticks := tickerWithDone(done, 250*time.Millisecond)

	go func() {
		for msg := range ticks {
			fmt.Println(msg)
		}
		fmt.Println("ticker goroutine exited")
	}()

	// Let it tick a few times
	time.Sleep(900 * time.Millisecond)
	// Signal stop
	close(done)

	// Give the goroutine a moment to exit
	time.Sleep(300 * time.Millisecond)
	fmt.Println()

	// ------------------------------------------------------------------
	// 5. CONTEXT CANCELLATION DEMO
	// ------------------------------------------------------------------
	fmt.Println("=== 5. Context cancellation ===")

	ctx, cancel := context.WithCancel(context.Background())
	go doWorkWithContext(ctx, 1)

	// Let the worker run briefly
	time.Sleep(700 * time.Millisecond)

	// Cancel the context
	fmt.Println("main: cancelling context")
	cancel()

	// Give worker time to print its shutdown message
	time.Sleep(300 * time.Millisecond)

	fmt.Println("\nChannel patterns demo complete.")
}
