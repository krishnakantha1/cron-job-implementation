package cron

import "time"

func sleeper(h *Handler, messageQueue chan struct{}) {

	if h.pq.IsEmpty() {
		return
	}

	//timeToWakeUp := h.pq.Peek().run_time.Sub(time.Now())
	timeToWakeUp := time.Until(h.pq.Peek().run_time)

	waiter := time.NewTimer(timeToWakeUp)

	select {
	case <-waiter.C:
		h.PopJob()
	case <-messageQueue:
		waiter.Stop()
		return
	}
}
