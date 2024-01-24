package main

import "fmt"

func main() {
	// declered slice
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{1: 12}

	// print slice
	fmt.Println(b == c)
	fmt.Println(b == d)
	fmt.Println(b[0])

	// create slice
	a := make([]int, 10)
	fmt.Println(a)

	// cut slice
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom"}
	users1 := initialUsers[1:3]
	fmt.Println(users1)

	// add element
	a = []int{1, 2, 3}
	a = append(a, 4, 5)
	fmt.Println(a)

	// remove element
	a = []int{1, 2, 3, 4, 5, 6, 7}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
	
	// copy elements
	j := []int{1, 2, 3}
	k := make([]int, 3, 3)
	n := copy(k, j)
	fmt.Println(n)
}
