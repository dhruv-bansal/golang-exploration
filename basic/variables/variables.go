package main

import "fmt"

var a string               // single declaration and by default zero value
var b, c string            // multi declaration and by default zero value
var d, e string = "d", "e" // multi var declaration and initialization

//f := 10 shorthand declaration is not allowed here

var g = 1 // inferred type outside the function using var keyword

func main() {
	f := 10 // this is called int inferred type and the inferred here is a int value
	fmt.Println(a, b, c, d, e, f, g)

	fmt.Println(len(e)) // length function of string
	fmt.Println("type of a is %T", a)
	fmt.Printf("type of f is %T\n", f)
	fmt.Printf("type of a and f are %T and %T", a, f)
}
