package main

import (
	"fmt"
)

func main() {
	<-myFunc()
}

func myFunc() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}
