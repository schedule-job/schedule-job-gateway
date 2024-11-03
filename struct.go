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

type Info struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Members     []string `json:"members"`
}

type Item struct {
	Name    string                 `json:"name"`
	Payload map[string]interface{} `json:"payload"`
}

type InsertItem struct {
	Info    Info `json:"info"`
	Action  Item `json:"action"`
	Trigger Item `json:"trigger"`
}
