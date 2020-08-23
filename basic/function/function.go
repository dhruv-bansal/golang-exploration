package main

import (
	"errors"
	"fmt"
)

func main() {

	result := sampleFunction(1, 2)
	fmt.Println(result)

	fmt.Println(functionReturnsError())

	// capture multiple return values
	p, err := functionReturnMultipleValues(4)
	fmt.Println(p, err)

	// write only value, it is not necessary to use those values
	_, err1 := functionReturnMultipleValues(6)
	fmt.Println(err1)
}

func sampleFunction(param1, param2 int) int {
	sumparam := param1 + param2
	return sumparam
}

// in go exceptions are not raised rather error and returns and
//it is the call of callie function what to do with those errors
func functionReturnsError() error {
	return errors.New("Some error occurred") // error going to print in the output
}

func functionReturnMultipleValues(param int) (int, error) {
	var result int
	var err error
	if param < 5 {
		result = param
	} else {
		err = errors.New("input greater the 5")
	}
	return result, err
}
