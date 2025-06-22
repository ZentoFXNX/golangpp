package main

import (
	"fmt"
	"time"
)

func sayMyNameAndDate() {
	name := "Lumoxyy"
	currentTime := time.Now()

	fmt.Printf("Привет, %s!\n", name)
	fmt.Printf("Сегодня: %s\n", currentTime.Format("02-01-2006"))
}

func main() {
	sayHelloWorld()
	sayMyNameAndDate()
}