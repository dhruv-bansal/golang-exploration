package main

import "fmt"

func main() {

	// init a map using make
	var map1 map[string]int = make(map[string]int)
	map1["1"] = 1
	fmt.Println(map1)

	// implicit declaration and initialization
	map2 := map[string]int{"foo": 1}
	fmt.Println(map2)
	fmt.Println(map2["foo"])

	map2["foo"] = 2
	fmt.Println(map2["foo"])

	// delete entry from map
	delete(map2, "foo")
	fmt.Println(map2)

}
