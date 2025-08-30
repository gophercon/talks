# Compile Time Errors My Belovéd
## Branden J Brown

Type errors from the compiler tell us when our code is obviously wrong, immediately, without ever having to run our code.
But it turns out there are techniques to turn "obviously wrong" into powerful proofs of correctness.
We can automatically detect code and tests that need to be updated as faraway definitions change.
We can have confidence that our programs run the same on every target, not just the ones we have test runners for.
Shall we check out static assertions in Go?

## Viewing the Slides

The presentation is written for [present](https://pkg.go.dev/golang.org/x/tools/cmd/present).
It is essentially a Markdown document with speaker notes on lines beginning with `:`. (The entire script is written in the speaker notes.)
To view the slides in their pretty format, clone this repo _recursively_ (the example code is in a submodule) and navigate to this directory, and then with the Go toolchain installed, run the command:

	go run golang.org/x/tools/cmd/present@latest -notes

This will launch an HTTP server serving the slides at http://127.0.0.1:3999/errors-my-beloved.slide.
You can press `n` with the slides open to see the speaker notes.

One of the slides mentions an invocation of the `xo` command to generate types from a database schema.
If you want to run that command, you'll first need to cd into errors-my-beloved/generated and run `sqlite3 bocchi.sqlite3 <./schema.sql`.
