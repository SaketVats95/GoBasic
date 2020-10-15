package main

import (
	"fmt"
)

func main() {

	for i := 5; i <= 10; i++ {
		fmt.Println("Saket ")

	}
	n1 := 10
	n2 := 20
	result, subs := calc(n1, n2)
	fmt.Println("Sum of Two numbers are :", result, " # And Differences are :", subs)
}

func calc(num1 int, num2 int) (int, int) {
	var result = num1 + num2
	var subs = num2 - num1
	return result, subs
}
