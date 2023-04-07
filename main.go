package main

import (
	"fmt"
	"time"

	c "github.com/krishnakantha1/planit-cron/cron"
)

func main() {
	h := c.Handler{}

	h.Init(10)

	//fmt.Println(time.Until(time.Date(2023, time.April, 7, 2, 8, 0, 0, time.UTC)))

	fmt.Println("second time insert")
	h.AddJob(2, time.Date(2023, time.April, 7, 14, 22, 5, 0, time.UTC))

	time.Sleep(5 * time.Second)

	fmt.Println("first time insert")
	h.AddJob(1, time.Date(2023, time.April, 7, 14, 22, 0, 0, time.UTC))

	for true {

	}
}
