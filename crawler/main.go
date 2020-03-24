package main

import (
	engine2 "gomodtest/crawler/engine"
	"gomodtest/crawler/zhenai/parser"
)

func main() {
	engine2.ConcurrentEngine{}.Run(engine2.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
