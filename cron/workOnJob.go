package cron

import (
	"fmt"
	"time"
)

func workOnJob(j job) {
	fmt.Println(j.id, j.run_time, time.Now())
}
