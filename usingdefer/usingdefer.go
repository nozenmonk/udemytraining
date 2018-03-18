package main

import "fmt"

func hello() {
	fmt.Printf("%s", "Hello")
}

func world() {
	fmt.Printf("%s", " World")
}

func newline() {
	fmt.Println("")
}

func main() {
	defer newline()
	defer world()
	hello() //Hello World \n because of defers
}
