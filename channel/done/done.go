package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {
	go func() {
		for n := range c {
			fmt.Printf("id:%d,chan:%c\n", id, n)
			go func() {
				done <- true
			}()
		}
	}()
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	c := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	doWorker(id, c.in, c.done)
	return c
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		//发送数据
		worker.in <- 'a' + i

	}

	for i, worker := range workers {
		//发送数据
		worker.in <- 'A' + i

	}

	//这里可能会出现的问题是：第一个<-worker.done还没收到;worker.in <- 'A' + i又开始发送数据了就会卡死
	//解决办法是把上面的返回done放在goroutine中
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
