package main

import (
	"fmt"
	"os"
)

func main() {

	for index, arguments := range os.Args {
		fmt.Print(index, " ", arguments)
	}
}
