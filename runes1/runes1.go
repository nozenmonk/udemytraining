package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println([]byte("hello"))

	for i := 50; i <= 4000; i++ {
		fmt.Println(i, " ", string(i), " ", []byte(string(i)))
	}
}
