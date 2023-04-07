package cron

import (
	"container/heap"
	"sync"
	"time"
)

type Handler struct {
	mutex       sync.Mutex
	pq          jobPQ
	messageChan chan struct{}
}

func (h *Handler) Init(size int) {
	//inititlize priorityQueue
	h.pq = make(jobPQ, 0, size)
	h.messageChan = make(chan struct{})
}

func (h *Handler) AddJob(id int64, date time.Time) {

	h.mutex.Lock()
	if !h.pq.IsEmpty() {
		h.messageChan <- struct{}{}
	}

	heap.Push(&(h.pq), job{id: id, run_time: date})
	h.mutex.Unlock()

	go sleeper(h, h.messageChan)
}

func (h *Handler) PopJob() {

	h.mutex.Lock()
	defer h.mutex.Unlock()

	for !h.pq.IsEmpty() && time.Until(h.pq.Peek().run_time) < 0 {
		val := heap.Pop(&(h.pq)).(job)

		workOnJob(val)
	}

	if !h.pq.IsEmpty() {
		go sleeper(h, h.messageChan)
	}
}
