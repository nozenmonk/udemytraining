package main

import "fmt"

/* return type is a func */
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//increment is a function which gets the value returned by wrapper
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
