package script

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

const CHANNEL_SIZE = 5

var currentJob []string

type Job struct {
	jobId int
	task  string
}

type Result struct {
	jobId  int
	result string
	err    error
}

//createWorkers will create worker pool and assign jobs to it
func CreateWorkers(task interface{}, noOfWorkers int) error {

	// interface type assert
	tempJob, ok := task.([]string)
	if !ok {
		return errors.New("jobs type mismatch")
	}
	currentJob = tempJob
	noOfJobs := len(currentJob)

	//buffered channels for jobs and and their results
	jobs := make(chan Job, CHANNEL_SIZE)
	results := make(chan Result, CHANNEL_SIZE)

	// channel for task completion
	done := make(chan bool)
	//allocate job to workers
	go allocate(noOfJobs, jobs)
	// print job results
	go result(done, results)
	//create pool of workers
	createWorkerPool(noOfWorkers, jobs, results)
	<-done
	return nil
}

// createWorkerPool will create parallel workers with given count
func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}
	wg.Wait()
	close(results)
}

// worker to work on the job inhand
func worker(wg *sync.WaitGroup, jobs chan Job, results chan Result) {
	for job := range jobs {
		result := getJobResult(job)
		results <- result
	}
	wg.Done()
}

// Allocate job to worker using jobs channel
func allocate(nofOfJobs int, jobs chan Job) {
	for jobId := 0; jobId < nofOfJobs; jobId++ {
		job := createJob(jobId)
		jobs <- job
	}
	close(jobs)
}

// result print data generated from worker to results channel
func result(done chan bool, results chan Result) {

	for result := range results {
		if err := result.err; err != nil {
			log.Printf("\n %s \n", err)
		} else {
			fmt.Printf("\n %s \n", result.result)
		}
	}
	done <- true
	// close(results)
}

//createJob creates the job to be completed by worker
func createJob(jobId int) Job {
	//create the job from given server URL
	job := Job{jobId, currentJob[jobId]}
	return job
}

//getJobResult calls the respective function to get the job response
func getJobResult(job Job) Result {
	// Get the url and hash of request
	data, err := GetHashFromURL(job.task)
	result := Result{job.jobId, data, err}
	return result
}
