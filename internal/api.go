package internal

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
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
			log.Fatalln(errResp.Error())
			return nil, errResp
		}

		defer resp.Body.Close()

		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			log.Fatalln(errRead.Error())
			return nil, errRead
		}

		return body, nil
	}

	log.Fatalln("no api")
	return nil, errors.New("no api")
}

func Post(path string, body io.Reader, urls []string) ([]byte, error) {
	for _, apiUrl := range urls {
		url := apiUrl + path

		resp, errResp := client.Post(url, "application/json", body)
		if errResp != nil {
			if errResp == context.DeadlineExceeded {
				continue
			}
			log.Fatalln(errResp.Error())
			return nil, errResp
		}

		defer resp.Body.Close()

		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			log.Fatalln(errRead.Error())
			return nil, errRead
		}

		return body, nil
	}

	log.Fatalln("no api")
	return nil, errors.New("no api")
}
