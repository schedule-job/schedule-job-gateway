package gateway

import "github.com/schedule-job/schedule-job-database/core"

type Agent struct {
	DB core.Database
}

type Batch struct {
	DB core.Database
}

type Job struct {
	DB core.Database
}

type Item struct {
	Info    core.Job     `json:"info"`
	Action  core.Action  `json:"action"`
	Trigger core.Trigger `json:"trigger"`
}
