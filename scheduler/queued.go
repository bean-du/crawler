package scheduler

import "github.com/bean-du/crawler/engine"

type QueuedSchedule struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedSchedule) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedSchedule) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedSchedule) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedSchedule) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 判断request列队和worker列队是否都有东西
			// 如果两个都有东西，那么就取一个request和一个worker
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				// worker 收到 request 之后将列队中的request第一个拿掉，同时拿掉worker列队的第一个
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
