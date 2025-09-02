

# Go Plays Nice With Your Computer - Race Detection and Freedom!
## Raghav Roy

### Abstract

Go is great for writing concurrent programs, but even if you write logically sound programs, you can still give way to data-races that are compiler or hardware dependent. What can you do to prevent them? How does Go help you detect races, and how do the latest changes to TSAN affect a Go dev?

In this demo, we explore the internals of Go's race detector, based on ThreadSanitizer (TSAN), and discover a fascinating connection between its core algorithm—Vector Clocks—and the principles of causality from Einstein's Special Relativity, as explored in the included papers by Einstein and F. Mattern. We will practically demonstrate the architectural leap from TSAN v2 (in Go 1.18) to TSAN v3 (in Go 1.19+), proving its massive improvements in scalability, speed, and memory efficiency.

### Links

*   [Link to the talk slides](https://github.com/RaghavRoy145/Gophercon2025/blob/main/Go%20Plays%20Nice%20With%20Your%20Computer%20Race%20Detection%20and%20Freedom!.pdf)
*   [Link to the References](https://github.com/RaghavRoy145/Gophercon2025/blob/main/References.pdf)
*   [Link to relevant papers](https://github.com/RaghavRoy145/Gophercon2025/tree/main/papers)
*   [Link to demo code](https://github.com/RaghavRoy145/Gophercon2025/tree/main/demo)
