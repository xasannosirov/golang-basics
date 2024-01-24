package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ExampleFunctionWithoutName() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	// example to anonym func
	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, src)

	fmt.Printf("Reverse string: %s. Result: %v.\n", src, src == test)
}

func Example() {
	// example to anonym func
	fn := func(a, b int) int { return a + b }

	// example to anonym func
	func(a, b int) {
		fmt.Println(a + b)
	}(12, 34)

	fmt.Println(fn(17, 15))
}

func externalFunction() func() {
	text := "TEXT"

	return func() {
		fmt.Println(text)
	}
}

func ExampleEnvironment() {
	//create func
	fn := externalFunction()
	fn()
}

func main() {
	ExampleFunctionWithoutName()
	Example()
	ExampleEnvironment()
}
