package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move() {
	fmt.Println(c.Model + " едет")
}

func (c Car) Stop() {
	fmt.Println(c.Model + " остановился")
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() {
	fmt.Println(b.Brand + " едет")
}

func (b Bicycle) Stop() {
	fmt.Println(b.Brand + " остановился")
}

func Drive(t Transport) {
	t.Move()
	t.Stop()
}

func main() {
	car := Car{Model: "Toyota"}
	bike := Bicycle{Brand: "Giant"}

	Drive(car)
	Drive(bike)
}