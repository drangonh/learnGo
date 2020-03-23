package parser

import (
	"fmt"
	"gomodtest/crawler/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1936803953" target="_blank">一直很安静</a>
const cityRe = `<a href=("http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//该文件是解析数据
func ParseCity(contexts []byte) engine.ParseResult {
	re, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}

	result := engine.ParseResult{}
	match := re.FindAllSubmatch(contexts, -1)
	for _, v := range match {
		//fmt.Printf("name:%s,url:%s\n", v[2], v[1])
		result.Items = append(result.Items, "user "+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(v[1]),
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, string(v[1]), string(v[2]))
				},
			})
	}
	fmt.Printf("%d\n", len(match))

	return result
}
