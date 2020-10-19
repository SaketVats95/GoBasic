package main

import "fmt"

func maxNumber(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num2
	} else {
		result = num1
	}
	return result
}

func main() {
	var a int = 100
	var b int = 200

	var ret = maxNumber(a, b)

	fmt.Printf("Max number is : %d\n", ret)
}
