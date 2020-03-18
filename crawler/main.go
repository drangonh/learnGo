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

	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	//<a href="http://www.zhenai.com/zhenghun/akesu" data-v-5e16505f>阿克苏</a>
	//<a href="http://www.zhenai.com/zhenghun/alashanmeng" data-v-5ef>阿拉善盟</a>
	//<a href="http://www.zhenai.com/zhenghun/aletai" data-v-5e16505f>阿勒泰</a>

	// <a target="_blank" href="http://www.zhenai.com/zhenghun/xuzhou">徐州征婚</a>
	// re,err:= regexp.Compile(`<a[^>]+href="http://www.zhenai.com/zhenghun/[^>]+">[^<]+</a>`)

	//<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/tacheng">塔城</a>
	//<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/aba">阿坝</a>
	re, err := regexp.Compile(`<a[^>]+href="http://www.zhenai.com/zhenghun/[^>]+"[^>]+>[^<]+</a>`)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n",contents)

	match := re.FindAll(contents, -1)
	for _, v := range match {
		fmt.Printf("%s\n", v)
	}
}
