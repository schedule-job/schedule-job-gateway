package internal

import (
	"encoding/json"
	"log"

	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func ToJson(data []byte) (interface{}, error) {
	var result map[string]interface{}

	var errUnmarshal = json.Unmarshal([]byte(string(data)), &result)

	if errUnmarshal != nil {
		err := schedule_errors.InternalServerError{Err: errUnmarshal}
		log.Fatalln(err.Error())
		return nil, &err
	}

	return result["data"], nil
}
