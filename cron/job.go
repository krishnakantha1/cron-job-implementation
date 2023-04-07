package cron

import (
	"time"
)

type job struct {
	id       int64
	run_time time.Time
}
