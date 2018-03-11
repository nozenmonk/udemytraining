package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	_, err := fmt.Println("Enter a meter value")
	fmt.Println(err)
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("The yards value is : ", yards)
}
