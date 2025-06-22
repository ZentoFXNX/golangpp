package main

import "fmt"

type Engine struct {
	Power    int
	Capacity float64
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine
}

func NewCar(brand string, model string, power int, capacity float64) Car {
	return Car{
		Brand:  brand,
		Model:  model,
		Engine: Engine{Power: power, Capacity: capacity},
	}
}

func PrintCar(c Car) {
	fmt.Printf("Марка: %s\nМодель: %s\nМощность двигателя: %d л.с.\nОбъем двигателя: %.1f л\n",
		c.Brand, c.Model, c.Engine.Power, c.Engine.Capacity)
}

func main() {
	car := NewCar("Toyota", "Camry", 200, 2.5)
	PrintCar(car)
}