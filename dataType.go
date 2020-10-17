package main

import "fmt"

func main() {
	var data float64
	data = 2000.000
	fmt.Println("Data : ", data)
	fmt.Printf("Data type of data is : %T\n", data) // output should be float64

	var a, b, c = 3, 4, "data" // Mixed Variable Declearation

	fmt.Printf("Data type of a : %T\n", a)

	fmt.Printf("Data type of b : %T\n", b)

	fmt.Printf("Data type of c : %T\n", c)
}
