package main

import (
	"fmt"
	"os"
	"testing"
)

type(
	Memoized func(int) int
)

var(
	fibMem = Memoize( fib )
)

func TestFib ( t *testing.T){

	fmt.Fprintf(os.Stdout, "%d\n", fibMem( 45 ))

}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	return fibMem( fib ( n ) )
}