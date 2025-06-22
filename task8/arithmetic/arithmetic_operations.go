package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\nАрифметические операции для чисел %.2f и %.2f:\n", a, b)
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
		fmt.Printf("%.2f %% %.2f = %.2f\n", a, b, float64(int(a)%int(b)))
	} else {
		fmt.Println("Деление и остаток на ноль невозможны!")
	}
}