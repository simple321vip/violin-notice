package queue

import "google.golang.org/genproto/googleapis/type/datetime"

type Job struct {
	JobId string
	datetime.DateTime
}

type TimeQueue struct {
	array *map[string][]Job
}

func NewQueue() *TimeQueue {
	m := make(map[string][]Job)
	return &TimeQueue{
		array: &m,
	}
}

func (q *TimeQueue) Push(job *Job) {
	jobList := q.array[job.JobId]
	if len(jobList) == 0 {
		jobList = append(jobList, job)
	}

}

func (q *TimeQueue) Pop() *Job {

	return &Job{}
}

func (q *TimeQueue) sort() *Job {

	return &Job{}
}
