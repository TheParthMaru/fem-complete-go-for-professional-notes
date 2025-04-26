package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello World")
	greetings("Parth")
	result := add(3, 4)
	fmt.Println(result)

	sum, product := sumAndProduct(3, 5)
	println(sum, product)
}

func greetings(name string) {
	fmt.Println("Hello", name)
}

func sumAndProduct(a, b int) (int, int) {
	return (a + b), (a * b)
}
