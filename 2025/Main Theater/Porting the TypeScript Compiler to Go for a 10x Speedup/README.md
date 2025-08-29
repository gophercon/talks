# Porting the TypeScript Compiler to Go for a 10x Speedup - Jake Bailey

[Link to talk](https://jakebailey.dev/talk-gophercon-2025)

From the beginning, the TypeScript compiler has been self-hosted, evolving alongside a growing ecosystem of millions of developers. As time went on, we faced challenges with the compiler's performance, largely inherent to the implementation language itself. Through experimentation and testing, we found Go to be an excellent language for our specific needs; a perfect porting language. In this talk, we will explore the process of porting the 150,000+ line TypeScript compiler and its 90,000+ tests to Go, the challenges we faced, lessons we learned, all leading to an overall 10x performance improvement over our previous implementation.
