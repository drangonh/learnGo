package main

import (
	"fmt"
)

func ss(i int) {
	fmt.Printf("我是第%d个函数\n", i)
}

func main() {
	for i := 1; i < 10; i++ {
		go func(i int) {
			fmt.Printf("我是第%d个函数\n", i)
			//for  {
			//	fmt.Printf("我是第%d个函数\n", i)
			//}
		}(i)
	}

	//加一个延迟是为了在第一个for结束之后main函数还未结束。
	//time.Sleep(time.Millisecond)
}
