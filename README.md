[![Go Report Card](https://goreportcard.com/badge/github.com/agirot/syncWorker)](https://goreportcard.com/badge/github.com/agirot/syncWorker)
# syncWorker
Simple program to parallelize exec command and redirect output with metadata.<br>
In my case a set of looong symfony (php framework) console command.

## Features
- Custom worker amount
- Centralization of workers output with metadata (start/finish time, total duration, worker id...)

## Overview

### config.json

```json
{
  "log_path": ".",
  "binary": "echo",
  "command": "/var/foo/bar command:test",
  "args": "--in=%v --out=%v",
  "args_to_replace": [
    [
      "foo",
      "bar"
    ],
    [
      "fromage",
      "baguette"
    ],
    [
      "pain",
      "béret"
    ]
  ]
}
```

### worker.log
```json
{
  "args_value": [
    "foo",
    "bar"
  ],
  "log_display": "/var/foo/bar command:test --in=foo --out=bar\n",
  "worker_id": 2,
  "args": "--in=foo --out=bar",
  "start": "2018-04-03T18:51:45.956470174+02:00",
  "finish": "2018-04-03T18:51:45.961112252+02:00",
  "total_time": "4.642078ms"
}
{
  "args_value": [
    "fromage",
    "baguette"
  ],
  "log_display": "/var/foo/bar command:test --in=fromage --out=baguette\n",
  "worker_id": 1,
  "args": "--in=fromage --out=baguette",
  "start": "2018-04-03T18:51:45.956527146+02:00",
  "finish": "2018-04-03T18:51:45.961341188+02:00",
  "total_time": "4.814042ms"
}
{
  "args_value": [
    "pain",
    "béret"
  ],
  "log_display": "/var/foo/bar command:test --in=pain --out=béret\n",
  "worker_id": 2,
  "args": "--in=pain --out=béret",
  "start": "2018-04-03T18:51:45.961451588+02:00",
  "finish": "2018-04-03T18:51:45.964618065+02:00",
  "total_time": "3.166477ms"
}
```
