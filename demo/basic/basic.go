package main

import (
	"fmt"
	"math"
)

func triangle() {
	var a, b int
	calcTriangle(a, b)
}

func calcTriangle(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	b := 1 << 5
	fmt.Println(b)

	a, b := 3, 4
	swap(&a, &b)

	fmt.Println(a, b, &a, &b)
}

func swap(a, b *int) {

	fmt.Println(a, b, *a, *b)

	a, b = b, a

	fmt.Println(a, b, *a, *b)
}
