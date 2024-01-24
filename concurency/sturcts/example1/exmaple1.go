package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	go myFunc(done)
	<-done
}

func myFunc(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}
