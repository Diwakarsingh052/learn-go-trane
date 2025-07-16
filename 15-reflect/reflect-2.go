package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type info struct {
	Data map[string]string
}

var jsonData = []byte(`{
		"name": "Elephant",
		"diet": "Herbivore",
		"size": "Large",
		"a_random_field": "Random Value",
		"1": 1
	}`)

func (i info) UnmarshalJSON(data []byte) error {
	rawMap := make(map[string]any)

	// Unmarshal the data from JSON into rawMap
	err := json.Unmarshal(data, &rawMap)

	if err != nil {
		return err // If there is an error, return it
	}

	i.Data = make(map[string]string)

	for key, value := range rawMap {
		val := reflect.ValueOf(value)

		// If the value is a string, store it in the Data map of the Animal struct.
		// If the value is not a string, it is ignored.
		if val.Kind() == reflect.String {
			i.Data[key] = val.String()
		}

		// If everything went well, return nil (i.e., no error)

	}
	return nil
}

func main() {

	var i info
	err := json.Unmarshal(jsonData, &i)
	if err != nil {
		panic(err)
	}
	fmt.Println(i.Data)
}
