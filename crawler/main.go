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
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code err", resp.StatusCode)
		return
	}

	//转换成utf-8
	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s",all)
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re, err := regexp.Compile(`<a[^>]+href="(http://www.zhenai.com/zhenghun/[^>]+)"[^>]+>([^<]+)</a>`)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n",contents)

	match := re.FindAllSubmatch(contents, -1)
	for _, v := range match {
		fmt.Printf("name:%s,url:%s\n", v[2], v[1])
	}
	fmt.Printf("%d", len(match))
}
