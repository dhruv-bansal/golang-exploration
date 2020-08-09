package main

import "fmt"

func main() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 1
	fmt.Println(arr, arr[2])
	//fmt.Println(arr[3]) // compile time array

	// implicit array declaration
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

}
