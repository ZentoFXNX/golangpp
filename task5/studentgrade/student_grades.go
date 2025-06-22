package main

import "fmt"

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func findGrade(grades map[string]int, name string) {
	grade, exists := grades[name]
	if exists {
		fmt.Printf("%s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден.\n", name)
	}
}

func deleteGrade(grades map[string]int, name string) {
	_, exists := grades[name]
	if exists {
		delete(grades, name)
		fmt.Printf("Оценка для %s удалена.\n", name)
	} else {
		fmt.Printf("Студент %s не найден, удаление невозможно.\n", name)
	}
}

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Alice", 90)
	addGrade(grades, "Bob", 85)
	addGrade(grades, "Charlie", 92)

	findGrade(grades, "Alice")
	findGrade(grades, "Diana")

	deleteGrade(grades, "Bob")

	fmt.Println("Итоговый список оценок:", grades)
}