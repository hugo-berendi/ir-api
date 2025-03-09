package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadDataFromFile() (Data, error) {
	var data Data

	// Open the file
	file, err := os.Open("data.json")
	if err != nil {
		return data, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return data, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse the JSON into the Data struct
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return data, nil
}
