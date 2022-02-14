package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	testingFunction := []string{"2", "3", "test", "7", "1.11", "false",
		"[{\"number\":1356,\"id\":\"1\",\"name\":\"test\"}]",
		"{\"number\":1356,\"id\":\"1\",\"name\":\"test\"}", "true", "21", "1.25", "OK"}

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
		} else if isJson(value) != nil {
			convertedList = append(convertedList, isJson(value))
		} else if isArray(value) != nil {
			convertedList = append(convertedList, isArray(value))
		} else {
			convertedList = append(convertedList, value)
		}
	}

	return convertedList
}

func isJson(str string) map[string]interface{} {
	var mapList map[string]interface{}
	err := json.Unmarshal([]byte(str), &mapList)
	if mapList != nil && err == nil {
		return mapList
	}
	return nil
}

func isArray(str string) []map[string]interface{} {
	var mapList []map[string]interface{}
	err := json.Unmarshal([]byte(str), &mapList)
	if mapList != nil && err == nil {
		return mapList
	}
	return nil
}
