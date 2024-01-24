package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ZeroDivisionError")
	} else {
		return a/b, nil
	}
}

func finish() {
	fmt.Println("Program has been finished")
}

func main() {
	// declered error
	err := errors.New("my error")
	fmt.Println(err)
	err2 := fmt.Errorf("%s", "my error2")
	fmt.Println(err2)
	// check error
	var a, b int
	fmt.Scan(&a, &b)
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// defer 
	defer finish()
	defer fmt.Println("Program has been started")
	fmt.Println("Program is working")
}
