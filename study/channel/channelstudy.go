/*
@Time : 2020/10/19 11:29 上午
@Author : mac
@File : channelstudy.go
@Software: GoLand
*/
package channel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
 * @Description goroutine和chanel的demo
 * @Param
 * @return
 **/
func ChanDemo() {
	var ch chan int
	fmt.Println(ch, &ch)

	c := newChan()
	go sendChan(c)
	go acceptChan(c)
	time.Sleep(time.Second)
}

/**
 * @Description  强制使用一个处理器。一个四核的处理器，就有4个处理器，能够同时并行处理4个任务。并行是同时处理多个任务，并发是同时管理多个任务。
 * @Param
 * @return
 **/
func singleThread() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("A:", i)
		}
		time.Sleep(time.Second * 5)
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("B:", i)
		}
		time.Sleep(time.Second * 5)
	}()
	wg.Wait()
}

func newChan() chan int {
	c := make(chan int, 3)
	return c
}

/**
 * @Author
 * @Description //TODO
 * @Date
 * @Param
 * @return
 **/
func acceptChan(c chan int) {
	for range c {
		fmt.Println(<-c)
	}

}

func sendChan(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

/**
 * @Description 协程用close关闭
 * @Param
 * @return
 **/
func printOne(cs chan int) {
	fmt.Println(1)
	cs <- 1
}

func printTwo(cs chan int) {
	<-cs
	fmt.Println(2)
	defer close(cs)
}

/**
 * @Description 同步输出
 * @Param
 * @return
 **/
func ChanDemo2() {
	cs := make(chan int)
	// 这里需要加关键字go,否则因为chan在写操作而卡死
	go printOne(cs)
	printTwo(cs)
}

func sendString(c chan string, string2 string) {
	go func() {
		c <- string2
	}()
}

func selectChan(ch1, ch2 chan string) {
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		}
	}
}

/**
 * @Description //select
 * @Param
 * @return
 **/
func ChanDemo3() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	sendString(ch1, "hello")
	sendString(ch2, "world")
	// 使用select前需要加go,否则也会死锁
	go selectChan(ch1, ch2)
	time.Sleep(time.Second)
}
