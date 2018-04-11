package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	first:= 0
	second:= 1
	fib:=0

	return func(next int) int {
		if next < 2{
			return next
		}
		fmt.Printf("i:%d fib=%d + %d", next, first, second)
		fib = first + second
		first = second
		second = fib
		fmt.Printf("= %d \n",fib)

		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		f(i)
	}
}
