package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
	fmt.Println("Goroutines have finished executing")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d - gorutina started work \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("%d - gorutine finishes work \n", id)
}
