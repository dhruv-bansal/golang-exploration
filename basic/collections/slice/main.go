package main

import "fmt"

func main() {

	// creating slice from array
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	slice := arr[:] // : represent from the beginning of the given collection to end of the collection
	fmt.Println(slice)

	arr[0] = 4
	fmt.Println(arr, slice) // slice is pointing to underline array

	// create slice
	// internal working is like Arraylist in Java
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	// create slice from slice
	slice3 := slice1[1:] // slice3 will be slice1 from the first index and proceeding to the end
	slice4 := slice1[:3]
	fmt.Println(slice3, slice4)
	slice1[2] = 43
	fmt.Println(slice1, slice3)
	slice3[1] = 52
	fmt.Println(slice1, slice3, slice4) // all these slices are referring to same memory location/array

}
