# Profiling Request Latency with Critical Path Analysis

[slides.pdf](./slides.pdf)

Go ships with great tools for diagnosing performance bottlenecks, with pprof’s CPU profiler being perhaps the most well-known and used tool.

However, when it comes to diagnosing request latency problems, CPU profiling is often insufficient, as Go request latency is usually dominated by Off-CPU waits for downstream services or databases.

This means that developers commonly rely on metrics, traces, or logs to break down their latency problems, requiring them to spend time, money, and even overhead on maintaining the desired level of observability. Despite these efforts, such instrumentation still leaves significant blind spots when it comes to mutex contention, channel bottlenecks, scheduling latency, backoff sleeps, and other Off-CPU waits, especially when they are the cause of high tail latencies.

So what if there was a better way to break down any Go latency issue to the code causing it, that requires zero instrumentation and comes with extremely low overhead?

If that sounds interesting, this talk is for you. You will learn about a novel implementation of critical path analysis to convert runtime execution tracing data into profiles and how you can use it to root cause even the most challenging tail latency problems with minimal effort.