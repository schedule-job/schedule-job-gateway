package gateway

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schedule-job/schedule-job-database/core"
	"github.com/schedule-job/schedule-job-gateway/internal"
)

type Agent struct {
	DB core.Database
}

func (a *Agent) getAgentUrls() []string {
	urls, error := a.DB.SelectAgentUrls()
	if error != nil {
		return []string{}
	}
	return urls
}

func (a *Agent) GetLogs(jobId, lastId string, limit int) ([]byte, error) {
	query := []string{}
	if lastId != "" {
		query = append(query, "lastId="+lastId)
	}
	if limit > 0 {
		query = append(query, "limit="+strconv.Itoa(limit))
	}

	path := fmt.Sprintf("/api/v1/request/%s/logs?%s", jobId, strings.Join(query, "&"))

	resp, err := internal.Get(path, a.getAgentUrls())

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return resp, nil
}

func (a *Agent) GetLog(jobId, id string) ([]byte, error) {
	path := fmt.Sprintf("/api/v1/request/%s/log/%s", jobId, id)

	resp, err := internal.Get(path, a.getAgentUrls())

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return resp, nil
}
