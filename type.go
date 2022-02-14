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
	var mapJson map[string]interface{}
	var mapList []map[string]interface{}

	for _, value := range args {
		if number, errNumber := strconv.Atoi(value); errNumber == nil {
			convertedList = append(convertedList, number)
		} else if float, errFloat := strconv.ParseFloat(value, 64); errFloat == nil {
			convertedList = append(convertedList, float)
		} else if boolean, errBoolean := strconv.ParseBool(value); errBoolean == nil {
			convertedList = append(convertedList, boolean)
		} else if errJson := json.Unmarshal([]byte(value), &mapJson); errJson == nil {
			convertedList = append(convertedList, mapJson)
		} else if errList := json.Unmarshal([]byte(value), &mapList); errList == nil {
			convertedList = append(convertedList, mapList)
		} else {
			convertedList = append(convertedList, value)
		}
	}

	return convertedList
}
