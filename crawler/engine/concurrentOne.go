package engine

import (
	"log"
)

type ConcurrentEngineOne struct {
	SchedulerOne SchedulerOne
	WorkerCount  int
}

type SchedulerOne interface {
	WorkerNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type WorkerNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngineOne) RunOne(seeds ...Request) {
	out := make(chan ParseResult)
	e.SchedulerOne.Run()

	//控制同时开放的goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorkerOne(e.SchedulerOne.WorkerChan(), out, e.SchedulerOne)
	}

	for _, r := range seeds {
		e.SchedulerOne.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("GOt it item %v", item)
		}

		for _, request := range result.Requests {
			e.SchedulerOne.Submit(request)
		}
	}
}

func createWorkerOne(in chan Request, out chan ParseResult, s WorkerNotifier) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
