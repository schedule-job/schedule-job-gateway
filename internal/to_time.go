package internal

import (
	"encoding/json"
	"log"
	"time"
)

func ToTime(data []byte) (*time.Time, error) {
	var result map[string]interface{}

	var errUnmarshal = json.Unmarshal([]byte(string(data)), &result)

	if errUnmarshal != nil {
		log.Fatalln(errUnmarshal.Error())
		return nil, errUnmarshal
	}

	layout := "2006-01-02T15:04:05Z"
	t, errParse := time.Parse(layout, result["data"].(string))

	if errParse != nil {
		log.Fatalln(errParse.Error())
		return nil, errParse
	}

	return &t, nil
}
