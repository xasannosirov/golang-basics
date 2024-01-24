package main

import "fmt"

func main() {
	// create a new channel
	channel := make(chan int, 5) 
	
	defer close(channel) 

	// print channel lenth and capacity
	fmt.Println(len(channel), cap(channel))

	// write data to channel
	channel <- 5
	channel <- 35
	channel <- 65

	// print channel lenth and capacity
	fmt.Println(len(channel), cap(channel))

	// read channel data
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// print channel lenth and capacity
	fmt.Println(len(channel), cap(channel))

}