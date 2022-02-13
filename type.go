package main

import (
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

	testing := [6]string{"2", "3", "test", "7", "1.11", "false"}

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
		} else {
			testAfter[i] = testing[i]
		}
	}

	for i := range testAfter {
		fmt.Println("value =", testAfter[i], "type =", reflect.TypeOf(testAfter[i]))
	}
}
