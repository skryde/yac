package main

import (
	"time"
)

type Cmd struct {
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type Cron struct {
	Command   Cmd           `json:"command"`
	TimeLapse time.Duration `json:"time_lapse"`
	TimeUnit  time.Duration `json:"time_unit"`
}

type Crons []Cron
