package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Course  int
	Average float64
}

func NewStudent(name string, age int, course int, average float64) Student {
	return Student{Name: name, Age: age, Course: course, Average: average}
}

func PrintStudent(s Student) {
	fmt.Printf("Имя: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n", s.Name, s.Age, s.Course, s.Average)
}

func UpdateAverage(s *Student, newAverage float64) {
	s.Average = newAverage
}

func main() {
	student := NewStudent("Alice", 20, 2, 4.5)
	PrintStudent(student)
	UpdateAverage(&student, 4.8)
	fmt.Println("\nПосле обновления среднего балла:")
	PrintStudent(student)
}