package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E

	radius := 5.0
	area := pi * radius * radius

	fmt.Println("Значение π:", pi)
	fmt.Println("Значение e:", e)
	fmt.Printf("Площадь круга с радиусом %.2f = %.2f\n", radius, area)
}