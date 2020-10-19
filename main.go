package main

import (
	"fmt"
	"gomodtest/algorithm"
)

func main() {
	fmt.Println(23)
	var b int
	b = algorithm.Gcd(4, 8)
	fmt.Println(b)
}
