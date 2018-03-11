package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a's value : ", a)
	fmt.Println("a's address : ", &a)
	fmt.Println("decimal hexa type")
	fmt.Printf(" %d %v %T\n", &a, &a, &a)
}
