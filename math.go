package main

import "fmt"

func main() {
	runProgram(112, 10)
}

func runProgram(a, b int) {
	fmt.Println(soma(a, b))
	fmt.Println(subtract(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(divide(a, b))
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
