package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		//延迟打印
		time.Sleep(time.Second * 5)
		//fmt.Printf("id:%d,chan:%c\n", id, n)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	//values收集数据
	var values []int
	n := 0
	//hasValue := false

	//在等待时间过去后，然后发送当前时间
	tm := time.After(10 * time.Second)

	//定时接收时间
	tk := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		//if hasValue {
		//	activeWorker = worker
		//}

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n = <-c1:
			//hasValue = true
			values = append(values, n)
		case n = <-c2:
			//hasValue = true
			values = append(values, n)

		case activeWorker <- activeValue:
			//hasValue = false
			//fmt.Println(values)
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("超时了")
		case <-tk:
			fmt.Printf("当前%T,values长度为%d\n", <-tk, len(values))
		case <-tm:
			fmt.Println("时间到10秒结束")
			return
		}
	}
}
