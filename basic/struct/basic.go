package main

import "fmt"

func main() {

	// struct can be at any level that make sense
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user // when it is initialized the each member in struct is initialize with zero values
	fmt.Println(u)

	// assigning values of struct
	u.ID = 1
	u.FirstName = "fn"
	u.LastName = "ln"
	fmt.Println(u)

	// short way to declaration of structs
	u2 := user{ID: 2,
		FirstName: "fnn",
		LastName:  "lnn", // need comma there or closing braces
	} // closing
	fmt.Println(u2)
}
