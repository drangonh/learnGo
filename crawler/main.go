package main

import (
	engine2 "gomodtest/crawler/engine"
	"gomodtest/crawler/scheduler"
	"gomodtest/crawler/zhenai/parser"
)

func main() {
	e := engine2.ConcurrentEngineOne{
		SchedulerOne: &scheduler.QueuedScheduler{},
		WorkerCount:  1,
	}
	e.RunOne(engine2.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	//单任务爬虫
	//engine2.Run(engine2.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})
}
