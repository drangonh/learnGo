package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 5, 2, 10, 3, 6}
	sort.Ints(a)
	for i, num := range a {
		fmt.Println(i, num)
	}
}
