package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Println(weekday)

	switch weekday {

	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Weekday")
	default:
		fmt.Println("No day")

	}
}
