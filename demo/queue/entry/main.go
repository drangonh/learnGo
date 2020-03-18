package main

import (
	"fmt"
	"gomodtest/demo/queue"
)

func main() {
	q := queue.Queue{2, 9, 4, 10, 30}
	q.Push(25)
	q.Pop()

	fmt.Println(q, q.IsEmpty())
}
