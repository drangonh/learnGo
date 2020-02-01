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

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) string {
	return poster.Post("https://www.imooc.com",
		map[string]string{
			"name": "hl",
		})
}

//接口组合只要包含Retriever，Poster中含有的方法即可
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(poster RetrieverPoster) string {
	poster.Post("i am a superman", map[string]string{
		"contents": "change old",
	})
	return poster.Get("www.imooc.com")
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

	s := &mock.Retriever{"hello"}
	fmt.Println(s)
	fmt.Printf("%T %v", s, s)
	fmt.Println()
	inspect(s)
	fmt.Println(session(s))
}

//r不是简单的值引用，r里面还有类型和值
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("i am mock.Retriever", v.Contents)
	case *mock.Retriever:
		fmt.Println("i am *mock.Retriever", v.Contents)
	case *real2.Retriever:
		fmt.Println("i am real2.Retriever", v.TimeOut, v.UseAgent)
	}
}
