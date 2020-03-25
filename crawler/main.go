package main

import (
	engine2 "gomodtest/crawler/engine"
	"gomodtest/crawler/scheduler"
	"gomodtest/crawler/zhenai/parser"
)

func main() {
	//并发版爬虫二
	//e := engine2.ConcurrentEngineOne{
	//	SchedulerOne: &scheduler.QueuedScheduler{},
	//	WorkerCount:  1,
	//}
	//e.RunOne(engine2.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	//并发版爬虫一
	//e := engine2.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount:  1,
	//}
	//e.Run(engine2.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	//单任务爬虫
	//engine2.Run(engine2.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	//只爬取一个城市的数据
	e := engine2.ConcurrentEngineOne{
		SchedulerOne: &scheduler.QueuedScheduler{},
		WorkerCount:  1,
	}
	e.RunOne(engine2.Request{
		Url:       "http://www.zhenai.com/zhenghun/luoyang",
		ParseFunc: parser.ParseCity,
	})
}
