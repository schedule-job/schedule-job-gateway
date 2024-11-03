package internal

import (
	"encoding/json"
	"log"
	"time"

	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func ToTime(data []byte) (*time.Time, error) {
	var result map[string]interface{}

	var errUnmarshal = json.Unmarshal([]byte(string(data)), &result)

	if errUnmarshal != nil {
		err := schedule_errors.InternalServerError{Err: errUnmarshal}
		log.Fatalln(err.Error())
		return nil, &err
	}

	layout := "2006-01-02T15:04:05Z"
	t, errParse := time.Parse(layout, result["data"].(string))

	if errParse != nil {
		err := schedule_errors.InternalServerError{Err: errParse}
		log.Fatalln(err.Error())
		return nil, &err
	}

	return &t, nil
}
