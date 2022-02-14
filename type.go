package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	testingFunction := []string{"2", "3", "test", "7", "1.11", "false", "{\"number\":1356,\"id\":\"1\",\"name\":\"test\"}"}

	for i := range listToTypedList(testingFunction) {
		fmt.Println("value =", listToTypedList(testingFunction)[i], ", type =", reflect.TypeOf(listToTypedList(testingFunction)[i]))
	}
}

func listToTypedList(args []string) []interface{} {

	var convertedList []interface{}

	for _, value := range args {
		if number, errNumber := strconv.Atoi(value); errNumber == nil {
			convertedList = append(convertedList, number)
		} else if float, errFloat := strconv.ParseFloat(value, 64); errFloat == nil {
			convertedList = append(convertedList, float)
		} else if boolean, errBoolean := strconv.ParseBool(value); errBoolean == nil {
			convertedList = append(convertedList, boolean)
		} else if isValidJSON(value) {
			var jsonMap map[string]interface{}
			data, err := getBytes(value)
			if err != nil {
				panic(err.Error())
			}
			data = data[4:]
			err = json.Unmarshal(data, &jsonMap)
			convertedList = append(convertedList, jsonMap)
		} else {
			convertedList = append(convertedList, value)
		}
	}

	return convertedList
}

func isValidJSON(i interface{}) bool {
	var str map[string]interface{}
	data, err := getBytes(i)
	if err != nil {
		panic(err.Error())
	}
	data = data[4:]
	err = json.Unmarshal(data, &str)
	return err == nil
}

func getBytes(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
