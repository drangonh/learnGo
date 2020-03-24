package scheduler

import "gomodtest/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	s.workerChan = r
}

func (s *SimpleScheduler) WorkerReady(c engine.Request) {

}

func (s *SimpleScheduler) Submit(c engine.Request) {
	// send request down to worker chan
	go func() {
		s.workerChan <- c
	}()
}
