package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		transform.NewReader()
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", all)
	} else {
		fmt.Println("status code err", resp.StatusCode)
		return
	}
}
