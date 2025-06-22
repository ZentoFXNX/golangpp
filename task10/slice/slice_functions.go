package main

import (
	"fmt"
	"sort"
)

func findElement(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func sortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func filterSlice(slice []int, condition func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 3, 7, 1}

	index := findElement(numbers, 8)
	fmt.Println("Индекс элемента 8:", index)

	sorted := sortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	filtered := filterSlice(numbers, func(n int) bool {
		return n > 4
	})
	fmt.Println("Элементы больше 4:", filtered)
}