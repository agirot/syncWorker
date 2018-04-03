package main

import (
	"sync"

	"github.com/agirot/syncWorker/config"
	"github.com/agirot/syncWorker/worker"
)

func init() {
	config.HydrateConfiguration()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(config.WorkerUnit)

	jobs := make(chan worker.Job, config.WorkerUnit)

	for w := 1; w <= config.WorkerUnit; w++ {
		go worker.Process(w, jobs, &wg)
	}

	for _, arg := range config.Config.ArgsValue {
		jobs <- worker.Job{ArgsValue: arg}
	}
	close(jobs)

	wg.Wait()
}
