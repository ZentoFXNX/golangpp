package main

import "fmt"

func changeByValue(n int) {
	n = n + 10
}

func changeByPointer(n *int) {
	*n = *n + 10
}

func main() {
	num1 := 50
	num2 := 50

	changeByValue(num1)
	fmt.Println("После передачи по значению:", num1)

	changeByPointer(&num2)
	fmt.Println("После передачи по указателю:", num2)
}