package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)
	for _, v := range f {
		w, ok := m[v]
		if ok {
			m[v] = w + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
