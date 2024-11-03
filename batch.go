package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	schedule_errors "github.com/schedule-job/schedule-job-errors"
	"github.com/schedule-job/schedule-job-gateway/internal"
)

func (b *Batch) getBatchUrls() []string {
	urls, error := b.DB.SelectBatchUrls()
	if error != nil {
		return []string{}
	}
	return urls
}

func (b *Batch) GetPreNextSchedule(name string, payload map[string]interface{}) (*time.Time, error) {
	path := fmt.Sprintf("/api/v1/schedule/pre-next/%s", name)
	body, errMarshal := json.Marshal(payload)

	if errMarshal != nil {
		err := schedule_errors.InvalidArgumentError{Param: "payload", Message: errMarshal.Error()}
		log.Fatalln(err.Error())
		return nil, &err
	}

	data, errReq := internal.Post(path, bytes.NewBuffer(body), b.getBatchUrls())

	if errReq != nil {
		return nil, errReq
	}

	time, errTime := internal.ToTime(data)

	if errTime != nil {
		return nil, errTime
	}

	return time, nil
}

func (b *Batch) GetNextSchedule(id string) (*time.Time, error) {
	path := fmt.Sprintf("/api/v1/schedule/next/%s", id)

	data, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		return nil, errReq
	}

	time, errTime := internal.ToTime(data)

	if errTime != nil {
		return nil, errTime
	}

	return time, nil
}

func (b *Batch) GetPreNextInfo(name string, payload map[string]interface{}) (interface{}, error) {
	path := fmt.Sprintf("/api/v1/request/pre-next/%s", name)
	body, errMarshal := json.Marshal(payload)

	if errMarshal != nil {
		err := schedule_errors.InvalidArgumentError{Param: "payload", Message: errMarshal.Error()}
		log.Fatalln(err.Error())
		return nil, &err
	}

	data, errReq := internal.Post(path, bytes.NewBuffer(body), b.getBatchUrls())

	if errReq != nil {
		return nil, errReq
	}

	json, errJson := internal.ToJson(data)

	if errJson != nil {
		return nil, errJson
	}

	return json, nil
}

func (b *Batch) GetNextInfo(id string) (interface{}, error) {
	path := fmt.Sprintf("/api/v1/request/next/%s", id)

	data, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		return nil, errReq
	}

	json, errJson := internal.ToJson(data)

	if errJson != nil {
		return nil, errJson
	}

	return json, nil
}

func (b *Batch) Progress() error {
	path := "/api/v1/progress"

	_, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		return errReq
	}

	return nil
}

func (b *Batch) ProgressOnce(id string) error {
	path := fmt.Sprintf("/api/v1/progress/%s", id)

	_, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		return errReq
	}

	return nil
}
