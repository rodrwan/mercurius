package server

import "fmt"

// Worker ...
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

// Start ...
func (w Worker) Start() {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				// Here we can put an asyncronous function
				fmt.Printf("\nWorker %d: started %s\n", w.id, job.Name)
				// Here we need to send this reponse to API
				body, endpoint := job.CategoryCall()
				fmt.Println(APICall(body, endpoint))
				fmt.Printf("worker%d: completed %s!\n", w.id, job.Name)
			case <-w.quitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

// Stop ...
func (w Worker) Stop() {
	go func() {
		w.quitChan <- true
	}()
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}
