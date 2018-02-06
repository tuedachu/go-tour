package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walkrecursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walkrecursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walkrecursive(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	Walkrecursive(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	v1, ok1 := <-ch1
	v2, ok2 := <-ch2
	for ok1 && ok2 {
		fmt.Println("ch1 = ", v1, " and ch2 =", v2)
		if v1 != v2 {
			return false
		}
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(10)))
}
