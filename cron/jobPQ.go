package cron

type jobPQ []job

func (jb jobPQ) Len() int {
	return len(jb)
}

func (jb jobPQ) Less(i, j int) bool {
	return jb[i].run_time.Before(jb[j].run_time)
}

func (jb jobPQ) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

func (jb *jobPQ) Push(job_instance any) {
	*jb = append(*jb, job_instance.(job))
}

func (jb *jobPQ) Pop() any {

	top_job := (*jb)[len(*jb)-1]

	*jb = (*jb)[:len(*jb)-1]

	return top_job
}

//custom helpers
func (jb jobPQ) IsEmpty() bool {
	return len(jb) == 0
}

func (jb jobPQ) Peek() job {
	if len(jb) == 0 {
		panic("no items to peak!")
	}

	return jb[0]
}
