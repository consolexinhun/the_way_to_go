package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan int, 1)
	quit := make(chan int, 1)

	go func() {
		for i := 0; i < 25; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took amount of time : %s\n", delta)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
