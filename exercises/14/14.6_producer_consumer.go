package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	go producer(0, 10, c)
	go consumer(c, done)

	<-done
}

func producer(start, step int, out chan int) {
	for i := 0; i < 10; i++ {
		out <- start
		start += step
	}
	close(out) // 关闭写通道
}

func consumer(in chan int, done chan bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}

	done <- true
}
