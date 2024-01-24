package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Second * 2)

	timer := time.After(time.Second)
	<-timer

	ticker := time.Tick(time.Second)
	count := 0

	for {
		<-ticker
		fmt.Println("another tick")
		count++
		if count == 3 {
			break
		}
	}
}
