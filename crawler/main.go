package main

import (
	engine2 "gomodtest/crawler/engine"
	"gomodtest/crawler/engine/scheduler"
	"gomodtest/crawler/zhenai/parser"
)

func main() {
	e := engine2.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 1,
	}
	e.Run(engine2.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	//单任务爬虫
	//engine2.Run(engine2.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})
}
