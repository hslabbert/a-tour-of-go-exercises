package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run(s string) (string, error) {
	if s != "mystring" {
		e := &MyError{
			time.Now(),
			"it didn't work",
		}
		return s, e
	}
	var e error
	return s, e
}

func main() {
	s, err := run("somestring")
	if err != nil {
		fmt.Println("We have an error")
		fmt.Println(s, err)
		return
	}
	fmt.Printf("Got string %s; this is fine\n", s)
}
