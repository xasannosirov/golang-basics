package main

import (
	"fmt"
	"time"
)

func main() {
	<-work()
}

func work() <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		stop := time.NewTimer(time.Second)

		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop()

		for {
			select {
			case <-stop.C:
				return
			case <-tick.C:
				fmt.Println("тик-так")
			}
		}
	}()

	return done
}
