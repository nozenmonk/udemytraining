package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
	)
	const (
		D = iota
		E
		F
	)
	fmt.Println("iota  set :", A)
	fmt.Println("iota :", B)
	fmt.Println("iota :", C)
	fmt.Println("iota reset :", D)
	fmt.Println("iota :", E)
	fmt.Println("iota :", F)
}
