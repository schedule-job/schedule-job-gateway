package internal

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

func Get(path string, urls []string) ([]byte, error) {
	for _, apiUrl := range urls {
		url := apiUrl + path

		resp, errResp := client.Get(url)
		if errResp != nil {
			if errResp == context.DeadlineExceeded {
				continue
			}
			err := schedule_errors.ConnectionError{Address: url, Reason: errResp.Error()}
			log.Fatalln(err.Error())
			return nil, &err
		}

		defer resp.Body.Close()

		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			err := schedule_errors.ConnectionError{Address: url, Reason: errRead.Error()}
			log.Fatalln(err.Error())
			return nil, &err
		}

		return body, nil
	}

	err := schedule_errors.InvalidArgumentError{Param: "urls", Message: "All servers are down."}
	log.Fatalln(err)
	return nil, &err
}

func Post(path string, body io.Reader, urls []string) ([]byte, error) {
	for _, apiUrl := range urls {
		url := apiUrl + path

		resp, errResp := client.Post(url, "application/json", body)
		if errResp != nil {
			if errResp == context.DeadlineExceeded {
				continue
			}
			err := schedule_errors.ConnectionError{Address: url, Reason: errResp.Error()}
			log.Fatalln(err.Error())
			return nil, &err
		}

		defer resp.Body.Close()

		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			err := schedule_errors.ConnectionError{Address: url, Reason: errRead.Error()}
			log.Fatalln(err.Error())
			return nil, &err
		}

		return body, nil
	}

	err := schedule_errors.InvalidArgumentError{Param: "urls", Message: "All servers are down."}
	log.Fatalln(err)
	return nil, &err
}
