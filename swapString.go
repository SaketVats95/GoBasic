package main

import "fmt"

func main() {
	a, b := swapString("Saket", "Vats")
	fmt.Printf("%s %s \n", a, b)
}

func swapString(a, b string) (string, string) {
	return b, a
}
