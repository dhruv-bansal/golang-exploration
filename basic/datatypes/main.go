package main

import "fmt"

func main() {

	// explicit declaration
	var i int
	i = 32 // explicit initialization
	fmt.Println(i)

	// explicit declaration and initialization
	var f float32 = 3.14
	fmt.Println(f)

	// implicit declaration with initialization
	title := "go data types"
	fmt.Println("Title", title)

	// implicit declaration of Boolean
	b := false
	fmt.Println("Boolean", b)

	// complex data type is built in GO
	c := complex(3, 4)
	fmt.Println("Complex Variable", c)

	// multi assignment in the go
	// two variables in the left and two functions in the right
	real, img := real(c), imag(c)
	fmt.Println("Real number and imaginary number", real, img)

}
