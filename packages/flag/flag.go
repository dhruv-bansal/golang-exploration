package main

import (
	"flag"
	"fmt"
)

// https://gobyexample.com/command-line-flags

func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// using var and &addressof operator
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// call before using the flag variables
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}