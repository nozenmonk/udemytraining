package main

import "fmt"

var x rune = 34

func main() {
	fmt.Printf(" %d %v %T\n", x, x, x)
	increment := func() rune {
		x++
		return x
	}
	fmt.Printf(" %d %v %T\b\a\n", increment(), x, x)
}
