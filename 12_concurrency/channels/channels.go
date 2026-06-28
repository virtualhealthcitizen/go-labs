// Runnable companion to ../02_channels_examples.md.
//
// Channels are Go's primary mechanism for communication between goroutines:
// "Don't communicate by sharing memory; share memory by communicating." This
// program walks the core channel patterns from the ground up.
//
//	go run ./12_concurrency/channels
package main

import (
	"fmt"
	"time"
)

// sendMessage demonstrates a send-only channel parameter (chan<- string): the
// compiler forbids receiving from it inside this function.
func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

// pingPong shows directional channels for input (<-chan) and output (chan<-).
func pingPong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- "pong to: " + msg
}

// worker reads jobs from a receive-only channel and writes results to a
// send-only channel, returning when jobs is closed and drained.
func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(50 * time.Millisecond) // simulate work
		results <- fmt.Sprintf("worker %d processed job %d", id, job)
	}
}

func main() {
	// 1. Unbuffered channel: a send blocks until another goroutine receives.
	fmt.Println("=== 1. Unbuffered channel (send/receive) ===")
	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "hello from goroutine"
	}()
	fmt.Println("received:", <-ch) // blocks until the goroutine sends
	fmt.Println()

	// 2. Buffered channel: sends don't block until the buffer is full.
	fmt.Println("=== 2. Buffered channel ===")
	buf := make(chan int, 3) // capacity 3
	buf <- 10
	buf <- 20
	buf <- 30
	fmt.Println("len(buf):", len(buf), "cap(buf):", cap(buf))
	fmt.Println(<-buf, <-buf, <-buf)
	fmt.Println()

	// 3. Directional channels enforce send-only / receive-only at compile time.
	fmt.Println("=== 3. Directional channels ===")
	pings := make(chan string)
	pongs := make(chan string)
	go sendMessage(pings, "ping")
	go pingPong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()

	// 4. Closing a channel + ranging over it. range stops once the channel is
	//    closed and drained.
	fmt.Println("=== 4. Closing channels and ranging ===")
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // signal: no more values will be sent
	}()
	for n := range numbers {
		fmt.Println("got:", n)
	}
	fmt.Println()

	// 5. The comma-ok receive: `v, ok := <-ch` reports whether the value is a
	//    real send (ok == true) or the zero value from a closed channel
	//    (ok == false). This is what `range` checks under the hood.
	fmt.Println("=== 5. Comma-ok receive on a closed channel ===")
	done := make(chan int, 1)
	done <- 42
	close(done)
	for {
		v, ok := <-done
		if !ok {
			fmt.Println("channel closed and drained; stop")
			break
		}
		fmt.Println("received value:", v)
	}
	fmt.Println()

	// 6. Worker pool: fan out jobs to N workers, fan in their results.
	fmt.Println("=== 6. Worker pool ===")
	jobs := make(chan int)
	results := make(chan string)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs) // no more jobs; workers' range loops will exit
	}()
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()

	// 7. select with timeout and a non-blocking default.
	fmt.Println("=== 7. select with timeout and default ===")
	slowChan := make(chan string)
	go func() {
		time.Sleep(300 * time.Millisecond)
		slowChan <- "finished slow operation"
	}()
	select {
	case v := <-slowChan:
		fmt.Println("received:", v)
	case <-time.After(150 * time.Millisecond):
		fmt.Println("timeout: slow operation took too long")
	}

	// Non-blocking send: default fires when no receiver is ready.
	nonBlocking := make(chan string)
	select {
	case nonBlocking <- "try send":
		fmt.Println("sent to nonBlocking")
	default:
		fmt.Println("send would block, did not send")
	}

	fmt.Println("\nChannel demo complete.")
}
