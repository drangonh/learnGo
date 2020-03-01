package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, done func()) {
	go func() {
		for n := range c {
			fmt.Printf("id:%d,chan:%c\n", id, n)
			go func() {
				done()
			}()
		}
	}()
}

type worker struct {
	in   chan int
	wg   *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	c := worker{
		in: make(chan int),
		wg: wg,
		done: func() {
			wg.Done()
		},
	}
	doWorker(id, c.in, c.done)
	return c
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//一共20个等待执行
	wg.Add(20)
	for i, worker := range workers {
		//发送数据
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		//发送数据
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
