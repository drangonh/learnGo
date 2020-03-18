package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
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
		//转换成utf-8
		e := determineEncoding(resp.Body)

		utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", all)
	} else {
		fmt.Println("status code err", resp.StatusCode)
		return
	}
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
