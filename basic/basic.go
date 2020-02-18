package main

import "math"

func triangle() {
	var a, b int
	calcTriangle(a, b)
}

func calcTriangle(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {

}
