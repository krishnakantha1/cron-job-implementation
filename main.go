package main

import (
	"fmt"
	"time"

	c "github.com/krishnakantha1/planit-cron/cron"
)

func main() {
	fmt.Println("service starting...")
	//initilize the handler to add jobs
	h := c.Handler{}

	h.Init(10)

	//Add jobs as you like. to better simulate real world ive put some delay when new jobs are added to the queue
	h.AddJob(2, time.Date(2023, time.April, 7, 14, 22, 5, 0, time.UTC))
	time.Sleep(5 * time.Second) //Delay the next entry
	h.AddJob(1, time.Date(2023, time.April, 7, 14, 22, 0, 0, time.UTC))

	//wait forever
	for true {
	}
}
