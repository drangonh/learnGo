package main

import (
	"gomodtest/functional/fib"
)

func main() {
	f := fib.Fibonacci()
	fib.PrintFileContents(f)
}
