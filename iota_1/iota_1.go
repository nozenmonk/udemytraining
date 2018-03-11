package main

import "fmt"

func main() {
	const (
		A = iota
		B = iota * 10
		C = iota
	)
	fmt.Println("iota :", A)
	fmt.Println("iota * 10 :", B)
	fmt.Println("iota :", C)
}
