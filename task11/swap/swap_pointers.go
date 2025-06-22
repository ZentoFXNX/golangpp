package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Println("До обмена: x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("После обмена: x =", x, "y =", y)
}