package main

import "fmt"

func main() {
	// create map
	thisMap := map[int]string{
		1: "one",
		2: "two",
		3: "tree",
	}

	// print map
	fmt.Printf("%T", thisMap)

	// delete all value in map
	for key := range thisMap {
		delete(thisMap, key)
	}
	fmt.Println(thisMap)

	// print all value in map
	for _, value := range thisMap {
		fmt.Println(value)
	}

	// print all value in map 
	for key := range thisMap {
		fmt.Println(thisMap[key])
	}

	// print map with key
	fmt.Println(thisMap[1])

	// check value in map
	if value, check := thisMap[1]; check {
		fmt.Println(value)
	}
	if value, check := thisMap[4]; check {
		fmt.Println(value)
	}

	// print lenth map
	fmt.Println(len(thisMap))
}
