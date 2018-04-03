package worker

import (
	"fmt"
	"github.com/agirot/syncWorker/config"
	"github.com/gin-gonic/gin/json"
	"os"
	"os/exec"
	"sync"
	"time"
)

// LogFileName name of generated log file
const LogFileName = "worker.log"

// Job store all data needed to exec and return a completed job
type Job struct {
	ArgsValue  []string  `json:"args_value"`
	Log        []byte    `json:"-"`
	LogDisplay string    `json:"log_display"`
	WorkerID   int       `json:"worker_id"`
	Args       string    `json:"args"`
	Start      time.Time `json:"start"`
	Finish     time.Time `json:"finish"`
	TotalTime  string    `json:"total_time"`
}

// Process contain all logical of worker
func Process(workerID int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		j.prepare(workerID)

		/* #nosec */
		cmd := exec.Command(config.Config.Binary, config.Config.Command, j.Args)
		b, err := cmd.CombinedOutput()
		/*cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()*/
		if err != nil {
			j.Log = []byte(err.Error())
		} else {
			j.Log = b
		}
		j.Finish = time.Now()
		if err := j.archiveWorkerLog(); err != nil {
			panic(err.Error())
		}
	}
}

func (j *Job) prepare(workerID int) {
	j.Start = time.Now()
	j.WorkerID = workerID
	j.Args = fmt.Sprintf(config.Config.Args, j.replaceArg()...)
}
func (j *Job) replaceArg() []interface{} {
	args := make([]interface{}, len(j.ArgsValue))
	for i, v := range j.ArgsValue {
		args[i] = v
	}

	return args
}

func (j *Job) archiveWorkerLog() error {
	f, err := os.OpenFile(config.Config.LogPath+"/"+LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	if _, err := f.Write(j.writeLog()); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (j *Job) writeLog() []byte {
	j.LogDisplay = string(j.Log[:])
	j.TotalTime = j.Finish.Sub(j.Start).String()

	b, err := json.Marshal(j)
	if err != nil {
		return []byte(err.Error())
	}
	return b
}
