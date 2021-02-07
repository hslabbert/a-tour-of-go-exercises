package main

import (
	"fmt"
)

// https://play.golang.org/p/LVOsv2_pKiR
func main() {
	args := [2]string{"valid2", "invalid2"}
	TestVars(args)
}

func TestVars(args [2]string) bool {
	for _, a := range args {
		if IsValidVariable(a) {
			fmt.Printf("Found one! %s", a)
			return true
		}
	}
	panic("No valid vars! Panicking!")
}

func IsValidVariable(arg string) bool {
	switch arg {
	case
		"valid1",
		"valid2",
		"valid3":
		return true
	}
	return false
}
