package main

import "fmt"

func main() {

	// explicit declaration
	var i int
	i = 32 // explicit initilization
	fmt.Println(i)

	// explicit declaration and initilization
	var f float32 = 3.14
	fmt.Println(f)

	// implicit declaration with initilization
	title := "go data types"
	fmt.Println("Title", title)

}
