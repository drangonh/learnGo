package main

import (
	"fmt"
)

func tryRecover() {
	//匿名函数怎么调用？ 需要在最后加()即可
	defer func() {
		r := recover()
		//r.(error)判断是否为err
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err.Error())
		} else {
			panic(err)
		}
	}()

	//b:=0
	//a := 5 / b
	//fmt.Println(a)

	//err := errors.New("this is an err")
	//panic(err)

	panic(123)
}

func main() {
	tryRecover()
}
