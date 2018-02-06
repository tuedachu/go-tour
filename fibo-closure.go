package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	secondlast := 0
	return func() int {
		if last == 0 {
			last = 1
			return 0
		} else {
			result := last + secondlast
			secondlast = last
			last = result
			return result
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
