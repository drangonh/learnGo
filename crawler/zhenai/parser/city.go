package parser

import (
	"fmt"
	"gomodtest/crawler/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1936803953" target="_blank">一直很安静</a>
var cityRe = regexp.MustCompile(`<a href=("http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(
	`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

//该文件是解析数据
func ParseCity(contexts []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}
	match := cityRe.FindAllSubmatch(contexts, -1)
	for _, v := range match {
		fmt.Printf("name:%s,url:%s\n", v[2], v[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(v[1]),
				ParseFunc: ProfileParser(string(v[2])),
			})
	}
	fmt.Printf("%d\n", len(match))

	match = cityUrlRe.FindAllSubmatch(
		contexts, -1)
	for _, m := range match {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			})
	}
	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
