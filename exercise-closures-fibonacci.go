package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev int
	var curr int
	var next int
	farnext := 1
	return func() int {
		prev = curr
		curr = next
		next = farnext
		farnext = curr + next
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
