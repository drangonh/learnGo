package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	go func() {
		for n := range c {
			fmt.Printf("id:%d,chan:%c\n", id, n)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		//发送数据
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()

	c := make(chan int, 3)

	worker(0, c)

	c <- 'a' + 2
	//发送完数据之后关闭chan，还需要在接收处判断是否还有数据要接受
	close(c)
	time.Sleep(time.Millisecond)
}
