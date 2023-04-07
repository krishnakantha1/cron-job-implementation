package cron

import (
	"time"
)

// Describes an individual job
type job struct {
	id       int64
	run_time time.Time
}
