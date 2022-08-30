package main

import (
	"fmt"
)

type Any interface{}
type FibonacciFunc func(Any) (Any, Any)

func main() {
	fibonacciFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}

	fibonacci := BuildLazyIntEvaluator(fibonacciFunc, []uint64{0, 1})

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fibo: %v\n", i, fibonacci())
	}
}

func BuildLazyEvaluator(fibonacciFunc FibonacciFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	go func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = fibonacciFunc(actState)
			retValChan <- retVal
		}
	}()
	retFunc := func() Any {
		return <-retValChan
	}
	return retFunc
}

func BuildLazyIntEvaluator(fibonacciFunc FibonacciFunc, initState Any) func() uint64 {
	f := BuildLazyEvaluator(fibonacciFunc, initState)
	return func() uint64 {
		return f().(uint64)
	}
}
