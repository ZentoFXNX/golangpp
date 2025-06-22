package main

import "fmt"

func main() {
	var names []string

	names = append(names, "Alice")
	names = append(names, "Bob")
	names = append(names, "Charlie")
	names = append(names, "Diana")

	fmt.Println("Список имен:")
	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)
	}
}