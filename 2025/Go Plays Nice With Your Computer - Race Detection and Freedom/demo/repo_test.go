// file: repro_test.go
package main_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Set a number higher than the TSAN v2 limit of 8128
const numGoroutines = 10000
const numMemoryGoroutines = 5000
// TestGoroutineLimit demonstrates the failure in TSAN v2.
// It uses t.Run with t.Parallel to create many goroutines that
// are all alive at the same time.
func TestGoroutineLimit(t *testing.T) {
	// Give the test a bit more time to run, just in case.
	t.Parallel()

	for i := 0; i < numGoroutines; i++ {
		i := i // capture loop variable
		t.Run(fmt.Sprintf("goroutine_%d", i), func(st *testing.T) {
			st.Parallel()
			// Sleep to ensure the goroutine stays alive long enough
			// for all others to be created, thus exceeding the limit.
			time.Sleep(1 * time.Second)
		})
	}
}

// BenchmarkGoroutineOverhead measures the performance of creating and
// synchronizing many goroutines with the race detector enabled.
func BenchmarkGoroutineOverhead(b *testing.B) {
	// We run the test for b.N iterations. Go's benchmark tool will
	// automatically choose a suitable value for b.N.
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for j := 0; j < numGoroutines; j++ {
			go func() {
				defer wg.Done()
				// Do a tiny bit of work to ensure the goroutine is
				// properly scheduled and tracked by the race detector.
				runtime.Gosched()
			}()
		}

		wg.Wait()
	}
}
// BenchmarkMemoryFootprint is designed to measure the steady-state memory
// overhead of the race detector. It keeps a fixed number of goroutines
// alive and busy, allowing us to profile the heap.
// This version uses a sync.Mutex for compatibility with all Go versions.
func BenchmarkMemoryFootprint(b *testing.B) {
	// This benchmark doesn't loop with b.N. We run it once for a fixed
	// duration to measure the memory of a stable system.
	b.ReportAllocs()

	var counter int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	done := make(chan struct{})

	wg.Add(numMemoryGoroutines)
	for i := 0; i < numMemoryGoroutines; i++ {
		go func() {
			defer wg.Done()
			// This loop keeps the goroutine alive and accessing shared
			// memory, forcing the race detector to maintain its state.
			for {
				select {
				case <-done:
					return
				default:
					// This lock-protected critical section is instrumented
					// by the race detector as a synchronization event.
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}
		}()
	}

	// Let the system run for 2 seconds to reach a stable state.
	time.Sleep(2 * time.Second)

	// Signal goroutines to stop and wait for them to exit.
	close(done)
	wg.Wait()
}
