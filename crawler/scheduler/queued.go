package scheduler

import "gomodtest/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		var activeRequest engine.Request
		var activeWorker chan engine.Request
		if len(requestQ) > 0 && len(workerQ) > 0 {
			activeRequest = requestQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case r := <-s.requestChan:
			requestQ = append(requestQ, r)
		case w := <-s.workerChan:
			workerQ = append(workerQ, w)
		case activeWorker <- activeRequest:
			workerQ = workerQ[1:]
			requestQ = requestQ[1:]
		}
	}()
}
