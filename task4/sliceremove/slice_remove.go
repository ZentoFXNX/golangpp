package main

import "fmt"

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("Ошибка: индекс вне диапазона")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Diana"}

	fmt.Println("Изначальный список:", names)

	names = removeElement(names, 1)

	fmt.Println("После удаления элемента:", names)
}