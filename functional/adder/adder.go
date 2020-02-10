package main

import (
	"fmt"
)

//闭包实现的函数编程
//这个闭包函数不仅返回了一个函数，还返回了对自由变量sum的引用，最后番薯返回了sum
func adder() func(i int) int {
	//sum是自由变量
	var sum = 0
	return func(i int) int {
		//i参数和返回的函数中的参数都可以看作是局部变量
		sum += i
		return sum
	}
}

//下面是正统的函数编程，没有状态，没有变量
type iAdder func(base int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		fmt.Println("iAdder", v)
		return base + v, adder2(base + v)
	}
}

func main() {
	a2 := adder2(10)
	a := adder()
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Println(a(i))
		fmt.Println(s)
	}

}
