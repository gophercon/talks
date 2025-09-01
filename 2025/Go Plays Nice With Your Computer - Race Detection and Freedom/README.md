

# Go Plays Nice With Your Computer - Race Detection and Freedom!
## Raghav Roy

### Abstract

Go is great for writing concurrent programs, but even if you write logically sound programs, you can still give way to data-races that are compiler or hardware dependent. What can you do to prevent them? How does Go help you detect races, and how do the latest changes to TSAN affect a Go dev?

In this demo, we'll explore the internals of Go's race detector, based on ThreadSanitizer (TSAN), and discover a fascinating connection between its core algorithm—Vector Clocks—and the principles of causality from Einstein's Special Relativity, as explored in the included papers by Einstein and F. Mattern. We will practically demonstrate the architectural leap from TSAN v2 (in Go 1.18) to TSAN v3 (in Go 1.19+), proving its massive improvements in scalability, speed, and memory efficiency.

### Prerequisites

*   A Linux or macOS environment.
*   `git` installed.
*   A shell like `bash` or `zsh`.

### Directory Contents

*   `demo/main.go`: A program to generate CPU profiles for analysis.
*   `demo/repo_test.go`: A Go test file to demonstrate the goroutine limit and run benchmarks.
*   `cpu18.prof` & `cpu19.prof`: Pre-generated profiles from `main.go` for direct viewing.
*   `papers/`: Contains foundational papers by Einstein and Mattern on relativity and virtual time.

---

## Part 1: Environment Setup

To accurately compare the two versions of the race detector, we need to be able to switch between Go 1.18 and Go 1.19. We will use `goenv` for this.

**1. Install `goenv`**

```bash
git clone https://github.com/syndbg/goenv.git ~/.goenv
```

**2. Configure Your Shell**

Add the following lines to the end of your shell configuration file (`~/.bashrc` for Bash, `~/.zshrc` for Zsh).

```bash
# Add these lines to the end of your ~/.bashrc or ~/.zshrc
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"
```
**Important:** The `eval "$(goenv init -)"` line should be near the end of your file to ensure it takes precedence.

**3. Reload Your Shell**

To apply the changes, either close and reopen your terminal or run:
```bash
exec $SHELL
```

**4. Install Go Versions**

We will install a late patch release of Go 1.18 and Go 1.19.
```bash
goenv install 1.18.1
goenv install 1.19.1 # Or a newer 1.19+ version
goenv rehash
```

---

## Part 2: Running the Demonstrations

Navigate into the Demo\_Code directory to run these commands.

### Demo 1: The 8k Goroutine Limit

This test will prove that TSAN v2 crashes when a high number of goroutines are spawned, while TSAN v3 handles it correctly.

**1. Run on Go 1.18 (Expect Failure)**
```bash
goenv local 1.18.16
go mod init tsan_demo
go mod tidy
go test -v -race -count=1 -run=TestGoroutineLimit .
```
**Expected Outcome:** The test will run for a few seconds and then fail with the error: `race: limit on 8128 simultaneously alive goroutines is exceeded, dying`.

**2. Run on Go 1.19 (Expect Success)**
```bash
goenv local 1.19.13
go test -v -race -count=1 -run=TestGoroutineLimit .
```
**Expected Outcome:** The test will run for a couple of seconds and then **PASS**. This demonstrates that the hard limit has been removed in TSAN v3.

### Demo 2: Performance & Memory Benchmarks

This benchmark from `repro_test.go` measures the speed and memory overhead of tracking thousands of goroutines.

*(Note: On macOS, you may need to use `gtime -v` instead of `/usr/bin/time -v`. Install with `brew install gnu-time`.)*

**1. Benchmark Go 1.18 (TSAN v2)**
```bash
goenv local 1.18.1
go test -bench=BenchmarkGoroutineOverhead -run=^$ -race -benchtime=10x
```
Note the `ns/op` value from the Go benchmark output and the `Maximum resident set size` from the `time` command's output.

**2. Benchmark Go 1.19 (TSAN v3)**
```bash
goenv local 1.19.1
go test -bench=BenchmarkGoroutineOverhead -run=^$ -race -benchtime=10x
```
Compare the results. You should see that TSAN v3 is several times faster (`ns/op` is much lower), though its peak memory usage may be higher due to its ability to run more concurrently.

### Demo 3: Generating Flame Graphs

This demo uses `main.go` to generate CPU profiles that we can visualize as cpu-graphs to see what's happening under the hood.

**1. Generate Profile for Go 1.18**
```bash
goenv local 1.18.1
go run -race main.go
mv cpu.prof cpu18.prof
```

**2. Generate Profile for Go 1.19**
```bash
goenv local 1.19.1
go run -race main.go
mv cpu.prof cpu19.prof
```

**3. View the Flame Graphs**
Open two separate terminal windows to launch the interactive profilers.

In terminal 1:
```bash
go tool pprof -http=:8080 cpu18.prof
```
In terminal 2:
```bash
go tool pprof -http=:8081 cpu19.prof
```
Now, open `http://localhost:8080` and `http://localhost:8081`.

---

## Part 3: Viewing Pre-Generated Profiles (Shortcut)

If you don't want to set up the environment, you can still view the most insightful results using the pre-generated profiles included in this repository.

**1. View the Go 1.18 Graph**
```bash
go tool pprof -http=:8080 cpu18.prof
```
**What to Look For:** Open the graph and notice the very wide bar for `__sanitizer_internal_sched_yield`. This shows the system was spending most of its time waiting on internal locks.

**2. View the Go 1.19 Graph**
```bash
go tool pprof -http=:8081 cpu19.prof
```
**What to Look For:** The `sched_yield` bar is gone. The new "hot spot" is a much narrower bar for `__tsan_go_atomic32_compare_exchange`. This visually proves the architectural shift from a slow, lock-contention model to a fast, lock-free atomic model.
