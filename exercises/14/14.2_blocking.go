package main

import (
	"fmt"
)

func main() {
	/* deadlock
	out := make(chan int)
	out <- 2
	go f1(out)
	*/

	/* solution1 begin*/
	// out := make(chan int, 1)
	// out <- 2
	// f1(out)
	/* solution1 end*/

	/* solution2 */
	out := make(chan int)
	go f1(out)
	out <- 2
}

func f1(in chan int) {
	fmt.Println(<-in)
}
