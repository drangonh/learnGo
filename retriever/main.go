package main

import (
	"fmt"
	"gomodtest/retriever/mock"
	real2 "gomodtest/retriever/real"
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
	r = real2.Retriever{}
	fmt.Println(download(r))
}
