package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 2

func main() {
	start := time.Now()

	runtime.GOMAXPROCS(NCPU)
	fmt.Println(CalculatePi(5000))

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}

func CalculatePi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < NCPU; k++ {
		go term(ch, k*n/NCPU, (k+1)*n/NCPU)
	}

	f := 0.0

	for k := 0; k < NCPU; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}
