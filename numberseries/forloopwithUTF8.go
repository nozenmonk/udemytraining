package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("decimal: %d   binary : %b   hexa : %x utf8 : %q \n", i, i, i,i)
	}
}
