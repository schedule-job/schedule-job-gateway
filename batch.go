package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/schedule-job/schedule-job-database/core"
	"github.com/schedule-job/schedule-job-gateway/internal"
)

type Batch struct {
	DB core.Database
}

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
		log.Fatalln(errMarshal.Error())
		return nil, errMarshal
	}

	data, errReq := internal.Post(path, bytes.NewBuffer(body), b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return nil, errReq
	}

	time, errTime := internal.ToTime(data)

	if errTime != nil {
		log.Fatalln(errTime.Error())
		return nil, errTime
	}

	return time, nil
}

func (b *Batch) GetNextSchedule(id string) (*time.Time, error) {
	path := fmt.Sprintf("/api/v1/schedule/next/%s", id)

	data, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return nil, errReq
	}

	time, errTime := internal.ToTime(data)

	if errTime != nil {
		log.Fatalln(errTime.Error())
		return nil, errTime
	}

	return time, nil
}

func (b *Batch) GetPreNextInfo(name string, payload map[string]interface{}) (interface{}, error) {
	path := fmt.Sprintf("/api/v1/request/pre-next/%s", name)
	body, errMarshal := json.Marshal(payload)

	if errMarshal != nil {
		log.Fatalln(errMarshal.Error())
		return nil, errMarshal
	}

	data, errReq := internal.Post(path, bytes.NewBuffer(body), b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return nil, errReq
	}

	json, errJson := internal.ToJson(data)

	if errJson != nil {
		log.Fatalln(errJson.Error())
		return nil, errJson
	}

	return json, nil
}

func (b *Batch) GetNextInfo(id string) (interface{}, error) {
	path := fmt.Sprintf("/api/v1/request/next/%s", id)

	data, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return nil, errReq
	}

	json, errJson := internal.ToJson(data)

	if errJson != nil {
		log.Fatalln(errJson.Error())
		return nil, errJson
	}

	return json, nil
}

func (b *Batch) Progress() error {
	path := "/api/v1/progress"

	_, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return errReq
	}

	return nil
}

func (b *Batch) ProgressOnce(id string) error {
	path := fmt.Sprintf("/api/v1/progress/%s", id)

	_, errReq := internal.Post(path, nil, b.getBatchUrls())

	if errReq != nil {
		log.Fatalln(errReq.Error())
		return errReq
	}

	return nil
}
