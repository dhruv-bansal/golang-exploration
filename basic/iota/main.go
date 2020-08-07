package main

import "fmt"

// constant block to define package level constants
// iota allows to evolve constants through out the constants block
// constant expression has to be evaluate at compile time
const (
	first  = iota + 6  // everytime iota is written value would be incremented by 1
	second = 2 << iota // constant expression i.e. maths over the iota
	third              // here expression defined above would be reused so iota value would be 2 and [2 << 2 == 8]
)

// iota resets in new constant block
const (
	firstnew  = iota // everytime iota is written value would be incremented by 1
	secondnew        // takes the value of above expression i.e. iota and value of iota = 1 here
	thirdnew         // takes the value of above expression i.e. iota and value of iota = 2 here
)

const pi = 3.14

func main() {

	fmt.Println(pi)
	fmt.Println(first, second, third)
	fmt.Println(firstnew, secondnew, thirdnew)
}
