package main

import (
	"fmt"
	"sync"
	"time"
)

//传统的通过共享内存实现同步，所以需要Mutex保护
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	//写数据的时候锁住，写完成之后解锁
	a.lock.Lock()
	//defer先进后出，所有defer在程序最后执行
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
