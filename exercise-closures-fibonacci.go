package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var curr int
	var prev int
	next := 1
	return func() int {
		oldnext := next
		prev = curr
		next = curr + next
		curr = oldnext
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
