package main

import (
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

// This number is high enough to create contention, but below the 8128 limit.
const numGoroutines = 5000

// worker is the function that each goroutine will run.
// It continuously locks a mutex and increments a counter until the done channel is closed.
func worker(wg *sync.WaitGroup, mu *sync.Mutex, counter *int64, done <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			// This is the hot path. Every lock/unlock is instrumented
			// by the race detector, generating a lot of profiling data.
			mu.Lock()
			*counter++
			mu.Unlock()
		}
	}
}

func main() {
	// --- Profiler Setup ---
	// Create a file to store the CPU profile.
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start the CPU profiler.
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	// Stop the profiler when main() exits.
	defer pprof.StopCPUProfile()

	// --- Workload Setup ---
	var (
		counter int64
		mu      sync.Mutex
		wg      sync.WaitGroup
		done    = make(chan struct{})
	)

	// Launch the workers.
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(&wg, &mu, &counter, done)
	}

	// Let the workers run for a fixed duration. This is a "controlled burn".
	println("Running workload for 3 seconds...")
	time.Sleep(3 * time.Second)

	// Stop the workers and wait for them to finish.
	println("Stopping workers...")
	close(done)
	wg.Wait()

	println("Done. Profile saved to cpu.prof")
}
