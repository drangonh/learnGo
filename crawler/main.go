package main

import (
	engine2 "gomodtest/crawler/engine"
	"gomodtest/crawler/zhenai/parser"
)

func main() {
	engine2.Run(engine2.Request{
		Url:       "http://album.zhenai.com/u/1774362825",
		ParseFunc: parser.ParseCityList,
	})
}
