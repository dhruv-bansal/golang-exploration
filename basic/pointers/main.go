package main

import "fmt"

func main() {

	// pointer data type - holders address of the location in memory that holds the location for us
	var name *string = new(string) // need to init the pointer location
	// default value of point is nil
	fmt.Println(name)

	// assign value by de-referencing the value
	*name = "Arthur"
	fmt.Println(name)  // prints address
	fmt.Println(*name) // prints value

	var someint int = 10
	//fmt.Print(someint + 1) // pointer arithmetic is not allowed

	// "address-of" operation
	someintPtr := &someint
	fmt.Println(someint, someintPtr)

	someint = 20
	fmt.Println(someint, someintPtr) // value changes but address doesn't

}
