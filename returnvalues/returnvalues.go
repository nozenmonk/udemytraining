package main

import "fmt"

func give3values() (x, y, z int) {
	return 10, 11, 12
}

func main() {
	/* error as b1 is not used again
	a1, _, b1 := give3values()
	fmt.Println(a1)
	*/

	//cannot do := twice for same variable

	//works as expected
	a2, _, b2 := give3values()
	fmt.Println(a2)
	fmt.Println(b2)

	//works as expected
	a3, b3, _ := give3values()
	fmt.Println(a3)
	fmt.Println(b3)

	//works as expected
	_, a4, _ := give3values()
	fmt.Println(a4)

	//gives error. expected at least one variable, valid variable
	//_, _, _ := give3values()

	//error case
	/*
		  a5 := give3values()
			fmt.Println(a5)
	*/
}
