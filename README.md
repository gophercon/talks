# GopherCon Talks

This repo contains slides and other presentation materials from [GopherCon](https://gophercon.com) talks over the years, submitted by the speakers themselves. Each year is organized into two sections, one for main theater talks and a second for lightning talks. The content is organized based on the agenda from that year.

If materials are available for a talk, the title will be a link to a sub-folder in this repo. See [CONTRIBUTING.md](./CONTRIBUTING.md) if you are a speaker and would like your content included.

<hr/>

## 2025

### Main Theater

#### Go's Next Frontier - Cameron Balahan

Go is more than just a programming language—it's an industry pioneer. From its revolutionary concurrency model to its leadership in open-source security, Go's enduring focus on modern software engineering compels us to not just keep pace with industry advancements but to actively define and lead them.

In this keynote, we’ll explore Go's legacy of leadership and how its core design principles make it well-positioned to lead us into the next frontier: the convergence of AI and programming languages. By examining what the Go team is working on now, we’ll also share a vision of Go’s future developer experience and what we’re most excited about in the AI era.

#### LLM Agents in Production: Why Go is the Missing Piece - Kartik Ramesh

You might have heard of this thing called AI agents that have been transforming our industry lately. While there exist a plethora of ways to prototype agents, operating these in production brings a new set of challenges. This talk demonstrates why Go is uniquely suited for building scalable, production-ready agent systems. We'll build an LLM agent from scratch without using any agent frameworks, scale our system out to a multi-agent architecture, and leverage Go's cloud-native ecosystem to operate agents at scale.

#### Goroutines and Cells: Lessons in Goal-Directed Systems - Carlisia Campos

This talk explores a new way to approach concurrency in Go by drawing inspiration from bioelectric signaling in cells. It contrasts how biological systems achieve scalable, decentralized coordination with traditional concurrency patterns such as worker pools and fixed-rate limiters, which often struggle to adapt to dynamic workloads.

Biological systems use graded communication, rhythmic synchronization, and localized interactions to coordinate without centralized control. By applying these principles to Go concurrency, developers can design adaptive, nature-inspired systems that improve efficiency, fault tolerance, and self-organization.

Through conceptual models and real-world analogies, this talk will illustrate:

- How graded signaling enables dynamic resource allocation, reducing bottlenecks in goroutine coordination
- How rhythmic synchronization helps orchestrate periodic tasks more efficiently than independent timers
- How decentralized interactions lead to emergent consensus without centralized control

Attendees will leave with new mental models for designing concurrent systems, gaining insights into how nature-inspired approaches can make Go applications more resilient and scalable.

#### Go: A "Tailor-Made" Language for AWS Lambda - Travis Bale

At Capital One, we're building scalable, high-performance services with Go and AWS Lambda.

This talk will walk through why Go feels like it was designed specifically to work well with AWS Lambda. Attendees will learn about AWS Lambda concepts such as runtimes, cold starts, and extensions, and how Go as a language pairs exceptionally well with them.

#### The Code You Reviewed is Not the Code You Built - Jess McClintock

Ken Thompson's "Reflections on trusting trust" highlights the potential disparity between source code and the final built product, cautioning developers to question their trust model right down to the compiler itself. This principle is crucial when dealing with intricate trust models, as code reviews alone can not provide sufficient assurance in the behaviors of the build artifact.

Earlier this year, a malicious typosquat of the boltdb package was identified on the Go module proxy. There were several things of note about this typosquat, most prominent being the way that the attackers modified the git tags to obfuscate the malicious behavior. Even a security-minded developer who reviewed the code present on GitHub could have been tricked into using this package. In a sense, there were two levels of deception used to ship this code - the typosquat itself and the decoy code present on GitHub. This attack exposes the real-world practicality of deceiving traditional source reviews, threatening the foundation of open source consumption.

#### My Protobuf Module is Faster than Yours (Because I Cheated) - Tom Lyons

When handling over 2 trillion messages every day, every ounce of performance matters. This talk outlines the (over) optimizations that went into making our Go protobuf module as efficient as possible.

#### Building a Decentralized Social Media App with Go and ATProto - Gautam Dey

Bluesky is an exciting new decentralized social media platform, powered by the ATProtocol (ATProto). What makes ATProto so revolutionary, and how can Go developers leverage it to build scalable social apps?

In this talk, we’ll explore ATProto's core components and their interactions. Then, we’ll put theory into practice by building an AppView—a key part of the ecosystem—using Go. Our project will be a photo-sharing app with likes and comments, demonstrating how Go developers can contribute to the growing ATProto ecosystem.

By the end, you'll have a strong foundation in ATProto and a roadmap for building your own decentralized social applications.

#### Analysis and Transformation Tools for Go Codebase Modernization - Alan Donovan

In this talk, Alan will report recent progress in analysis and refactoring tools for Go. The talk will cover:

- The principles of the Go analysis framework, which powers batch tools like cmd/vet and interactive tools like gopls, and enables problems to be precisely reported and fixed in real time;
- a discussion of recent tools to "modernize" your codebase so that it uses the most efficient, clear, and idiomatic features of the newest Go releases; and
  a demonstration of how our algorithm for precise inlining of function calls is at the heart of a powerful tool for codebase migration.

#### Plowing Through Data: Building Flexible Pipelines with Go - Mindy Ratcliff

In modern agriculture, data flows in from everywhere—fields, labs, applications—and needs to be processed, transformed, and delivered to a wide range of destinations. Using Go’s parametric polymorphism, we have been able to build quickly with scalability and maintainability.

This talk will cover practical insights into building flexible pipelines with Go—and maybe even inspire you to plow through your own data challenges with a bit more ease.

#### Understanding Escape Analysis - Bill Kennedy

???

#### Scaling LLMs with Go: Production Patterns for Handling Millions of AI Requests - John Wang

At [Assembled](https://www.assembled.com/), we serve tens of millions of LLM requests, powering customer support AI for companies like Canva, Etsy, and Patreon. While most teams build LLM infrastructure in Python, we found Go's type system, concurrency features, and interfaces solve many production challenges more elegantly.

This talk shares concrete patterns from our experience deploying Go-based LLM systems at scale: from handling structured outputs to parallel RAG pipelines. You'll learn practical approaches for building performant, reliable AI applications leveraging Go's strengths.

#### Porting the TypeScript Compiler to Go for a 10x Speedup - Jake Bailey

From the beginning, the TypeScript compiler has been self-hosted, evolving alongside a growing ecosystem of millions of developers. As time went on, we faced challenges with the compiler's performance, largely inherent to the implementation language itself. Through experimentation and testing, we found Go to be an excellent language for our specific needs; a perfect porting language. In this talk, we will explore the process of porting the 150,000+ line TypeScript compiler and its 90,000+ tests to Go, the challenges we faced, lessons we learned, all leading to an overall 10x performance improvement over our previous implementation.

#### From Chaos to Cohesion: A Community of Practice Story - Moieed Ahmed and Seth Lumnah

An exploration of the challenges of managing languages in a large enterprise and the role that communities of practice play in providing standardization from a diverse set of opinions.

#### An Operating System in Go - Patricio Whittingslow

We’ll answer the age-old question — “Is Go a systems programming language?” — by showcasing TinyGo’s rise as the unexpected contender in embedded systems, upending assumptions about a space long thought settled.

#### Go’s Trace Tooling and Concurrency - William Kennedy

In this talk, Bill will share how to use Go’s trace tooling to examine a Go program's performance. Along the way, he will live-code a Go program, adding different levels of concurrency to see how the performance changes.

Throughout this talk, we will cover:

- Go’s trace tooling
- Goroutines and Concurrency semantics
- Performance exploration

#### Profiling Request Latency with Critical Path Analysis - Felix Geisendörfer

Go ships with great tools for diagnosing performance bottlenecks, with pprof’s CPU profiler being perhaps the most well-known and used tool.

However, when it comes to diagnosing request latency problems, CPU profiling is often insufficient, as Go request latency is usually dominated by Off-CPU waits for downstream services or databases.

This means that developers commonly rely on metrics, traces, or logs to break down their latency problems, requiring them to spend time, money, and even overhead on maintaining the desired level of observability. Despite these efforts, such instrumentation still leaves significant blind spots when it comes to mutex contention, channel bottlenecks, scheduling latency, backoff sleeps, and other Off-CPU waits, especially when they are the cause of high tail latencies.

So what if there was a better way to break down any Go latency issue to the code causing it, that requires zero instrumentation and comes with extremely low overhead?

If that sounds interesting, this talk is for you. You will learn about a novel implementation of critical path analysis to convert runtime execution tracing data into profiles and how you can use it to root cause even the most challenging tail latency problems with minimal effort.

#### Invisible Insight: Strategies for Auto-Instrumenting Go Applications Without Code Changes - Hannah Kim

Instrumenting Go applications without modifying source code has long been a challenge due to the language’s static nature and lack of dynamic code modification capabilities. This talk explores and compares ready-to-use, open-source solutions that enable auto-instrumentation at compile-time and runtime, analyzing their trade-offs in performance, stability, and security. We will cover leading approaches, including eBPF-based solutions, compile-time instrumentation, and runtime techniques such as shared library injection and binary trampoline methods, to assess their feasibility and limitations.

A major focus will be on how community collaboration in OpenTelemetry has made auto-instrumentation a pragmatic reality for Go applications. This session with explore ongoing Go runtime developments that could enable new patterns for auto-instrumentation, including improvements in function call interception, runtime hooks, and tracing APIs (golang/go#63185).

Attendees will gain benchmark-driven insights into these techniques, helping them evaluate the right auto-instrumentation approach for their production environments by balancing observability, performance overhead, and ease of deployment.

Whether you’re an SRE, performance engineer, or Go developer, this talk will provide actionable insights and production-ready open-source tools to integrate into your Go applications today.

#### Next-Gen AI Tooling with MCP Servers in Go - Katie Hockman

Step into this new era of AI-powered tools with the Model Context Protocol (MCP) server! This session will teach you how to expose your APIs directly into the IDE of your users so they can interact with your systems using natural language. The talk will center around a step-by-step demo where we'll write an MCP server together using the brand new github.com/modelcontextprotocol/go-sdk library. By the end of this session, you’ll be equipped with the knowledge to build your own MCP server to support your AI-enabled users.

#### Supercharging ML Pipelines with Go - Viadehi Thete

Building scalable and efficient machine learning pipelines often requires overcoming bottlenecks in data transfer, feature retrieval, and orchestration. This talk showcases how Go was leveraged to operationalize an ML model service, transforming it into a high-performance, real-time system. By utilizing Go’s powerful concurrency model, shared memory for inter-process communication, and efficient queuing mechanisms, we reduced inference times from hours to just 10-15 minutes.

We’ll explore the architecture that decouples Go’s operational responsibilities—such as feature retrieval, queuing, and shared memory management—from the ML model’s ranking tasks. Attendees will learn how shared memory was used to transfer millions of scores efficiently, how Go’s interfaces enabled rapid prototyping, and how its strong typing and safety ensured robust system design. Real-world benchmarks, implementation details, and lessons learned will provide actionable insights for engineers tackling similar challenges in high-performance computing and distributed systems.

Whether you’re a software engineer, system architect, or distributed systems enthusiast, this talk will demonstrate how Go can be the backbone of scalable and efficient ML inference pipelines.

#### Go Faster: Integrating CUDA in Go for GPU Acceleration - Sam Burns

In the era of machine learning and artificial intelligence, GPU acceleration has become indispensable for handling compute-intensive tasks efficiently. This talk explores how to integrate CUDA - NVIDIA's powerful parallel computing platform - with Go, to leverage GPU power in your applications. We'll start with an introduction to GPU computing and the CUDA programming model, explaining why and when to use GPU acceleration, particularly in the context of AI and ML. Finally, we'll guide you through setting up your development environment to work with Go and CUDA.

Using Go's cgo, we will demonstrate:

- How to call CUDA functions from Go;
- Managing memory between the CPU and GPU;
- Debugging and performance monitoring.

We will also cover common patterns that will help your applications run smoothly. Through practical examples, you will learn how to implement and optimize GPU-accelerated computations in Go. By the end of this talk, you'll be equipped to start incorporating GPU acceleration into your Go projects, achieving significant performance improvements for compute-intensive tasks, including those in machine learning and artificial intelligence.

#### Go Plays Nice With Your Computer - Race Detection and Freedom! - Raghav Roy

Go is great for writing concurrent programs, but even if you write logically sound programs, you can still give way to data races that are compiler- or hardware-dependent. What can you do to prevent them? How does Go help you detect races, and how do the latest changes to TSAN affect a Go dev? Also, does Einstein have anything to do with this?

#### Advancing Go Garbage Collection with Green Tea - Michael Knyszek

Memory latency and bandwidth are becoming increasingly constrained, and these trends are at odds with most of today's garbage collection algorithms, including Go's. In this talk, Michael will dive deep into Green Tea, a new parallel mark algorithm to accelerate Go's garbage collector through improved cache locality and the use of modern SIMD hardware. The talk will also cover what this new algorithm means for the performance of Go programs, along with suggestions on how advanced Gophers can capitalize on these improvements.

#### The Go Cryptography State of the Union - Filippo Valsorda

I know I say this every year, but Go 1.24 has been an exciting cycle for Go cryptography.

A new maintainer, post-quantum crypto, native FIPS 140, a crypto/rsa rewrite, moar systematic testing, x/crypto packages in the standard library, faster AES-CTR, more testable ECDSA... even a rare backwards compatibility break: rejecting RSA keys so small they are trivially broken. And then an external audit, and Encrypted Client Hello, and crypto/rand vDSO that never fails, and more.

This is your yearly appointment to learn about the progress of Go’s cryptography libraries and the techniques we use to maintain them.

#### AI and Go: Opportunities and Challenges - Johnny Boursiquot (host)

As artificial intelligence continues to reshape the practice of software engineering, Go developers face both exciting opportunities and unique challenges. This panel will explore how AI is influencing the Go ecosystem—from building AI-powered systems in Go to integrating with popular AI services. Panelists will discuss Go's role in AI-powered software and infrastructure as well as community efforts to bridge Go and AI. Attendees will gain insights into how Go continues to remain relevant and competitive now and in the AI future.

Panelists: David Soria Parra, Gari Singh, Ian Cottrell, Jaana Dogan

<hr/>

### Lightning Talks

#### Whodunit? Writing Murder Mysteries with LLMs - James Heller, Apartment 304

Can a large language model craft a brand-new murder mystery on demand? Explore a mystery-writing AI and the tips and tricks that help it generate coherent, logic-driven mysteries for players to solve. The system writes backwards from the solution, breaks the story into structured steps, and uses Temporal to manage retries when APIs fail. The result is a low-tech game powered by cutting-edge tools.

#### Standardizing Feature Flagging with OpenFeature - Sahid Velji, Capital One

Feature flags are essential for modern software delivery -- enabling safer rollouts, faster feedback, and more control. In this talk, you’ll learn about OpenFeature, a CNCF project that brings a standard, pluggable approach to feature flags across languages and platforms. You’ll see how easy it is to use with your favorite feature flag management system, and get a peek at the OpenFeature SDK in action.

#### Why Go Rocks for Building a Lua Interpreter - Roxy Light, 256 Lights

Many Go programs are web services or clients for web APIs. But Go can do more than that! The same features that make Go productive for writing backend systems also make Go ideal for writing a programming language interpreter. In this talk, Roxy will share her experience with writing a Lua interpreter in Go as part of the new zb build tool. You’ll learn how the interpreter is structured and what it took to adapt a mid-sized C library into Go.

#### Apps Without an Operating System?! - Andrew Williams, Apptrix.ai

The popularity of Go in unusual places has been shown by the TinyGo project and the recent bare metal support discussion backed by the Tamago implementation. What people might not realize is that you can compile and run graphical applications at this hardware level as well. In this Lightning Talk, Andrew will show how a Fyne GUI app can run directly from a QEMU EFI boot, removing all the requirements of an operating system.

#### Nil Today, Outage Tomorrow - Mukul Mantosh, JetBrains

We’ve all been there… you deploy in the afternoon and then stroll into work the next day to find everything is broken in production. You sit down, start debugging and checking the logs, and then you find the dreaded nil exception. Come along to my talk to learn how you can use your GoLand IDE to spot nil dereference errors before your code hits production. Avoid panics, embrace safety, and sleep well!

#### Go with the (Work)Flow: AI-Driven Code Generation for Workflows - Samantha Coyle, Diagrid

Discover how we leveraged AI to transform workflow diagrams into Dapr Workflow application code. This lightning talk will explore our implementation using an Intermediate Representation (IR) and graph visitor pattern to support AI-driven code generation, offering insights into building durable, scalable cloud-native workflow applications with Go.

#### Go Static Analysis with the Single Static Assignment - Takuya Ueda, newmo, Inc.

Single Static Assignment (SSA) form makes it easier to explicitly analyze operations on the same value over a control flow graph. Because SSA represents a program's control flow as a graph, general graph algorithms can be applied. In this talk, Takuya will demonstrate static analyses that he has implemented himself—such as checking whether a function argument can be nil, or whether cleanup operations are always performed—using the `golang.org/x/tools/go/ssa` package.

#### [Compile Time Errors My Belovéd - Branden J Brown](/2025/Compile%20Time%20Errors%20My%20Belovéd)

Type errors from the compiler tell us when our code is obviously wrong, immediately, without ever having to run our code. But it turns out there are techniques to turn "obviously wrong" into powerful proofs of correctness. We can automatically detect code and tests that need to be updated as faraway definitions change. We can have confidence that our programs run the same on every target, not just the ones we have test runners for. Shall we check out static assertions in Go?

#### Hurtling into the AI Coding Era - Michael Richman, Bitly

Join Michael as he takes you through the trials and tribulations of rolling out AI Coding Tools at a mid-size Eng Org, His Lightning Talk will touch on evaluating and choosing toolsets, reactions and resistance, best practices, as well as successes and failures experienced.

#### Not Your Parent’s Editor: A Gopher’s Guide to Neovim - Bethany Janos, GitHub

As Go continues to evolve, so do the tools we use to write it. This talk demonstrates how Neovim provides a powerful, flexible environment specifically tailored for Go development. Drawing from years of terminal-based development experience, this talk showcases tools that simplify configuration, which essential plugins enable IDE-like functionality for Go, and tips on persisting your setup across machines. Whether you're curious about terminal-based editing or looking to level up your existing Neovim setup, you'll leave with practical techniques to enhance your Go development workflow.

As a software engineer who has been working on large-scale Go services for 7 years, Bethany has learned a lot about refining her tooling to maximize productivity while maintaining a lightweight, keyboard-centric workflow. Her journey with Neovim has evolved alongside her Go development experience, resulting in a highly customized environment that balances simplicity with powerful language-specific features.

#### Git Server from Scratch in Go - Sergii Iefremov, Axios Media

In this talk, Sergii will dive into internals of Git protocol and his experience implementing Git SSH server in Go using excellent gliderlabs/ssh library. This specifically covers internal workings of `git clone` operation and steps Git normally does during its `git-upload-pack` operation. He’ll show some simple Go code that receives SSH requests and walks Git commit graph to select objects to send back to the client.

#### To Trie or Not to Trie: Typeahead Completion Using Redis - Garrett Denis, Skool

Ditch Postgres LIKE queries for a Redis-backed trie to supercharge @mentions. In this talk, you'll learn how we moved @mentions off Postgres into Redis, built a prefix tree for typeahead completion, and cut latency by 80% under heavy load.

#### A Small Update on TinyGo - GopherCon 2025 Edition - Ron Evans, The Hybrid Group

Come along on a lightning-speed tour of the latest and greatest happenings in the world of TinyGo.
