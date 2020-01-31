package main

import (
	"fmt"
	"gomodtest/downhtml/testing"
)

//func getRetriever() testing.Retriever {
//这里的返回类型可以是struct也可以是interface，只要这个interface中的方法在return的struct都有就可以
func getRetriever() retriever {
	return testing.Retriever{}
}

//返回一个可以Get的retriever,接口的出现是因为强类型语言中在使用变量的时候不关心它是属于哪个struct，只关心它有相关的方法就可以
type retriever interface {
	Get(url string) string
	//Post()
}

func main() {
	//其实main中不关心retriever是testing的还是infra的。main只需要拿到retriever的Get
	//retriever := getRetriever()
	//fmt.Println(retriever.Get("https://www.imooc.com"))

	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
