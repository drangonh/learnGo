package main

import (
	"fmt"
	"runtime"
	"time"
)

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second) //柱塞进程
	}
	stopCh <- true

}

func main() {
	//beego.Run();
	fmt.Printf(runtime.GOARCH)
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)
	m, _ := time.Parse("2006-01-01", "2018-05-01")
	fmt.Println(m)

	arr := [3]int{1, 2, 3}
	//遍历数组，打印值
	for index, res := range arr {
		println(index, res)
	}

	//下面的for会执行一个无限循环直到满足某一case然后退出循环
	for {
		select {

		case s := <-ch:
			fmt.Println("Receive", s)
		case c = <-ch:
			fmt.Println("Recvice", c)
			fmt.Println("channel")

			//_下划线在import中表示执行该模块该包中所有的init()函数，之后不能用模块名调用该模块的方法
			//下划线在代码中表示忽略这个变量
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
