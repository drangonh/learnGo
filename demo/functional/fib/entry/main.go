package main

import (
	"gomodtest/demo/functional/fib"
)

func main() {
	f := fib.Fibonacci()
	fib.PrintFileContents(f)
}
