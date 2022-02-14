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

	primes := []string{"2", "3", "test", "7", "11", "false"}

	test1, err2 := json.Marshal(primes)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		//fmt.Println("here")
		//fmt.Println(&test1[0])
	}

	var test []interface{}

	error2 := json.Unmarshal(test1, &test)

	if error2 != nil {
		fmt.Println(error2)
	}

	//fmt.Println(test)
	//for i := range test {
	//	fmt.Println(test[i])
	//}

	testing := [7]string{"2", "3", "test", "7", "1.11", "false", "{\"id\":\"1\"}"}

	var testAfter [len(testing)]interface{}

	for i := range testing {
		value := testing[i]
		number, errNumber := strconv.Atoi(value)
		float, errFloat := strconv.ParseFloat(value, 64)
		boolean, errBoolean := strconv.ParseBool(value)
		if errNumber == nil {
			testAfter[i] = number
		} else if errFloat == nil {
			testAfter[i] = float
		} else if errBoolean == nil {
			testAfter[i] = boolean
		} else if isValidJSON(value) {
			var jsonMap map[string]interface{}
			data, err := getBytes(value)
			if err != nil {
				panic(err.Error())
			}
			data = data[4:]
			err = json.Unmarshal(data, &jsonMap)
			testAfter[i] = jsonMap
		} else {
			testAfter[i] = testing[i]
		}
	}

	//for i := range testAfter {
	//	fmt.Println("value =", testAfter[i], ", type =", reflect.TypeOf(testAfter[i]))
	//}

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
