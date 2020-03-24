package parser

import (
	"fmt"
	"gomodtest/crawler/engine"
	"regexp"
)

const cityListRe = `<a[^>]+href="(http://www.zhenai.com/zhenghun/[^>]+)"[^>]+>([^<]+)</a>`

//该文件是解析数据
func ParseCityList(contexts []byte) engine.ParseResult {
	re, err := regexp.Compile(cityListRe)
	if err != nil {
		panic(err)
	}

	result := engine.ParseResult{}
	//获取城市的数组
	match := re.FindAllSubmatch(contexts, -1)

	for _, v := range match {
		//fmt.Printf("name:%s,url:%s\n", v[2], v[1])
		result.Items = append(result.Items, "city "+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(v[1]),
				ParseFunc: ParseCity,
			})
	}
	fmt.Printf("%d\n", len(match))

	return result
}
