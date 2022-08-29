package main

import "fmt"

func main() {
	c := make(chan int, 1)
	quit := make(chan int, 1)
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
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
