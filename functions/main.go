package main

import "fmt"

// this func find fibonacci numbers
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	var f1, f2 int = 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1
}

// this func print give arguments to func
func myPrint(a ...interface{}) {
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
}

// this func example to scope
func ExampleScope() {
	var v int = 1

	{
		var v string = "2"
		fmt.Println(v)
	}

	fmt.Println(v)
}

func main() {
	fmt.Println(fibonacci(4))
	myPrint(1, 2, 3, 4, 5)
	ExampleScope()
}
