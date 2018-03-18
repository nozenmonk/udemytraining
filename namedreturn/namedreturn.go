package main

import "fmt"

func main() {
	fmt.Println(greetPerson("John", "Mccain"))
}

func greetPerson(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
