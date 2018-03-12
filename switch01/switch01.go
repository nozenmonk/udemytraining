package main

import "fmt"

func main() {
	x := "yes"

	switch {
	case x == "no":
		fmt.Println("x is no")
	case x == "yes":
		fmt.Println("x is yes")

	}
}
