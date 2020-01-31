package main

import (
	"fmt"
	"gomodtest/retriever/mock"
	real2 "gomodtest/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"wo shi shi"}
	inspect(r)
	r = &real2.Retriever{
		TimeOut:  time.Minute,
		UseAgent: "5 year",
	}
	inspect(r)
	//fmt.Println(download(r))

	//通过type assertion来判断interface的类型
	//v, ok := r.(mock.Retriever)
	v, ok := r.(*real2.Retriever)
	if ok {
		fmt.Println(v.TimeOut)
	} else {
		fmt.Println("i am a mock.Retriever")
	}
}

//r不是简单的值引用，r里面还有类型和值
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("i am mock.Retriever", v.Contents)
	case *real2.Retriever:
		fmt.Println("i am real2.Retriever", v.TimeOut, v.UseAgent)
	}
}
