package util

import (
	"encoding/json"
	"io/ioutil"
)

func ReadData(fileName string) (map[string]any, error) {

	if fileName == "" {
		return nil, nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
