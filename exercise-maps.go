package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sf := strings.Fields(s)
	for _, w := range sf {
		_, present := m[w]
		if !present {
			m[w] = 0
		}
		m[w]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
