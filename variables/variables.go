package main

import "fmt"

var g string = "cowboy"

func main() {
	a := 10
	b := 4
	c := 'h'
	d := 'c'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(g)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf(" %v \n", d)

	s := true
	fmt.Printf(" %T %v\n", s, s)

	var h int
	var j rune
	var k float64
	var m bool
	fmt.Printf(" %T %v\n", h, h)
	fmt.Printf(" %T %v\n", j, j)
	fmt.Printf(" %T %v\n", k, k)
	fmt.Printf(" %T %v\n", m, m)
}
