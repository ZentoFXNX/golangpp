package main

import "fmt"

func main() {
	var year int

	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	// Логическое условие для проверки високосного года
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d - високосный год.\n", year)
	} else {
		fmt.Printf("%d - не високосный год.\n", year)
	}
}