package main

import "fmt"

func main() {

	// we dont have to initialize them at the time of declaration
	// value has to be confirmed at compile time
	const pi = 3.145
	fmt.Println(pi)

	// pi = 1.4 // reassigning constant not allowed

	//implicitly typed constants
	const intConstant = 3
	fmt.Println(intConstant + 3) //implicitly determines the type of constant is int

	fmt.Println(intConstant + 3.1) //implicitly determines the type of constant is float

	// if we explicitly specify the data type of constant then adding floating point would not be possible
	const constant int = 3
	fmt.Println(constant + 3) //explicitly determines the type of constant is int

	// fmt.Println(constant + 3.1) // compilation error

	// we have to do explicit conversion of int to float32
	fmt.Println(float32(constant) + 3.1)
}
