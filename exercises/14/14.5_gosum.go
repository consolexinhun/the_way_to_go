package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(x, y int, c chan int) {
		c <- x + y
	}(12, 13, c)

	fmt.Println(<-c)
}
