package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	fmt.Println("KB : ", KB)
	fmt.Println("MB : ", MB)
	fmt.Println("GB : ", GB)
}
