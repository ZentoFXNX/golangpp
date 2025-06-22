package main

import "fmt"

type Student struct {
	Name    string
	Birth   int
	Average float64
}

func (s Student) Age(currentYear int) int {
	return currentYear - s.Birth
}

func (s Student) Status() string {
	if s.Average >= 4.5 {
		return "Отличник"
	} else if s.Average >= 3.0 {
		return "Хорошист"
	}
	return "Троечник"
}

func (s Student) Speak() {
	fmt.Println("Привет от " + s.Name)
}

func main() {
	alice := Student{Name: "Alice", Birth: 2004, Average: 4.7}
	alice.Speak()
	fmt.Println("Возраст:", alice.Age(2025))
	fmt.Println("Статус:", alice.Status())
}