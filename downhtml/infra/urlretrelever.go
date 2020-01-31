package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Retriever struct {
}

//(r Retriever)r可以省略因为这里r用不到
func (Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	return string(bytes)
}

func (Retriever) Post() {
	fmt.Printf("测试")
}
