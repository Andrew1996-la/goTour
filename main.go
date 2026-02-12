package main

import (
	"sync"
)

type Job struct {
	index int
	value int
}

func RunWorkers(jobs []int, workerCount int) []int {
	result := make([]int, len(jobs))
	wg := &sync.WaitGroup{}
	jobChan := make(chan Job)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for job := range jobChan {
				result[job.index] = job.value * job.value
			}
		}()
	}

	for i, v := range jobs {
		jobChan <- Job{
			index: i,
			value: v,
		}
	}

	close(jobChan)
	wg.Wait()

	return result
}

func main() {

	RunWorkers([]int{1, 2, 3}, 2)
}
