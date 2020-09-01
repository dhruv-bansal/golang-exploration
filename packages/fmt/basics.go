package main

import "fmt"

func main() {

	var name string
	var name1 string
	var alphabet_count int

	fmt.Scan(&name)
	fmt.Scan(&name1)
	fmt.Println(name, name1)
	// Calling Sscanln() function
	fmt.Sscanln("", &name, &alphabet_count)

	// Printing the elements of the string
	fmt.Printf("name: %s, alphabet_count: %d", name, alphabet_count)
}
