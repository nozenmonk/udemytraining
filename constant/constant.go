package main

import "fmt"

const p int = 42

//const f := 3   error : must declare type for const

func main() {
	var q string = "Death and Taxes"
	const PI float64 = 3.14

	fmt.Println(" p :", p)
	fmt.Println(" q :", q)

	q = "liva da revolution"
	fmt.Println(" q :", q)

	/* error string set to int value
	q = 32
	fmt.Println(" q :", q)
	*/

	fmt.Println(" PI :", PI)
}
