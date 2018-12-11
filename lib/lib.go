package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseConfig() (map[string]map[string]string, error) {
	var result map[string]map[string]string
	data, err := ioutil.ReadFile("config.conf")
	if err != nil {
		return nil, err
	}

	// json to map
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
