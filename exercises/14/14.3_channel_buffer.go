package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received:", x)
	}()

	fmt.Println("sending :", 10)

	c <- 10
	fmt.Println("sent :", 10)
}

/* Output
sending 10
sent 10

因为 main() 直接结束了
*/
