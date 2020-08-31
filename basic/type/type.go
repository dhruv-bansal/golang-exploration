package main

import "fmt"

func main() {

	type Celsius float64 // creating the custom type using default type float
	type IDNum int32     // creating the custom type using default type int

	var temp Celsius = 1.1
	fmt.Println(temp)

	var id IDNum = 123
	fmt.Println(id)
}
