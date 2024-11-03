package internal

import (
	"encoding/json"
	"log"
)

func ToJson(data []byte) (interface{}, error) {
	var result map[string]interface{}

	var errUnmarshal = json.Unmarshal([]byte(string(data)), &result)

	if errUnmarshal != nil {
		log.Fatalln(errUnmarshal.Error())
		return nil, errUnmarshal
	}

	return result["data"], nil
}
