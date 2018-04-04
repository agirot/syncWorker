package config

import (
	"encoding/json"
	"os"

	flag "github.com/spf13/pflag"
)

// Config content current loaded config
var Config ConfigurationFile

// WorkerUnit Nb of concurrent goroutine
var WorkerUnit int

// ConfigurationFile store all data of config.json
type ConfigurationFile struct {
	Binary    string     `json:"binary"`
	Command   string     `json:"command"`
	Args      string     `json:"args"`
	ArgsValue [][]string `json:"args_to_replace"`
	LogPath   string     `json:"log_path"`
}

func init() {
	flag.IntVar(&WorkerUnit, "worker", 2, "Number of worker")
	flag.Parse()
}

// HydrateConfiguration check and store config file
func HydrateConfiguration() {
	file, err := os.Open("config.json")
	if err != nil {
		panic("config.json not found")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic("Failed parsing json config: " + err.Error())
	}
}
