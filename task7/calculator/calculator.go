package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Результат: %.2f\n", num1/num2)
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	default:
		fmt.Println("Ошибка: неизвестный оператор!")
	}
}