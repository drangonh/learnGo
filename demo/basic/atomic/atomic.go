package main

import (
	"fmt"
	"time"
)

//这个例子会造成数据冲突，原因是数据在读的时候，也在执行写的操作
type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
