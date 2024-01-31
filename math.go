package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(subtract(112, 10))
	fmt.Println(multiply(112, 10))
	fmt.Println(divide(112, 10))
}

func soma(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}
