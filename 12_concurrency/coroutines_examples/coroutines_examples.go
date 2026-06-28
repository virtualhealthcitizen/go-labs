// goroutines_examples.go
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// worker simulates doing some work, then sends a result on the channel.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Simulate work taking a random time
		sleepMs := rand.IntN(500) + 100 // 100–600 ms

		time.Sleep(time.Duration(sleepMs) * time.Millisecond)

		results <- fmt.Sprintf("Worker %d processed job %d in %dms", id, job, sleepMs)
	}
}

// simpleGoroutine demonstrates a basic goroutine that runs concurrently with main.
func simpleGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("[simpleGoroutine] tick", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	// --------------------------------------------------------------------
	// 1. Basic goroutine with WaitGroup
	// --------------------------------------------------------------------
	fmt.Println("=== 1. Basic goroutine with WaitGroup ===")

	var wg sync.WaitGroup
	wg.Add(1)
	go simpleGoroutine(&wg)

	for i := 1; i <= 3; i++ {
		fmt.Println("[main] doing work", i)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println()

	// --------------------------------------------------------------------
	// 2. Channels — unbuffered and buffered
	// --------------------------------------------------------------------
	fmt.Println("=== 2. Channels: unbuffered and buffered ===")

	// Unbuffered channel
	unbuffered := make(chan string)

	go func() {
		defer close(unbuffered)
		unbuffered <- "message via unbuffered channel"
	}()

	msg := <-unbuffered
	fmt.Println("Received:", msg)

	// Buffered channel
	buffered := make(chan int, 3)
	buffered <- 10
	buffered <- 20
	buffered <- 30
	close(buffered)

	for v := range buffered {
		fmt.Println("From buffered channel:", v)
	}
	fmt.Println()

	// --------------------------------------------------------------------
	// 3. Worker pool with goroutines
	// --------------------------------------------------------------------
	fmt.Println("=== 3. Worker pool with goroutines ===")

	const numWorkers = 3
	const numJobs = 7

	jobs := make(chan int)
	results := make(chan string)

	wg = sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
	fmt.Println()

	// --------------------------------------------------------------------
	// 4. Using select with timeout
	// --------------------------------------------------------------------
	fmt.Println("=== 4. select + timeout ===")

	timeoutChan := make(chan string)

	go func() {
		delayMs := rand.IntN(800)
		delay := time.Duration(delayMs) * time.Millisecond // FIXED
		time.Sleep(delay)
		timeoutChan <- fmt.Sprintf("Finished slow operation in %v", delay)
	}()

	select {
	case msg := <-timeoutChan:
		fmt.Println("Result:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timed out waiting for slow operation")
	}

	fmt.Println("\nAll goroutine demos complete.")
}
