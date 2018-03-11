package main

import "fmt"

func changeValue(x int) {
	x = 0
}

func changeValueAddr(y *int) {
	*y = 0
}

func main() {
	x := 5

	changeValue(x)
	fmt.Println(x)
	changeValueAddr(&x)
	fmt.Println(x)
}
