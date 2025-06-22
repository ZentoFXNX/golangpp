package main

import "fmt"

func main() {
	var intVar int = 100
	var int8Var int8 = -8
	var int16Var int16 = 1600
	var int32Var int32 = -32000
	var int64Var int64 = 6400000000

	var uintVar uint = 100
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65000
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	var float32Var float32 = 3.14
	var float64Var float64 = 2.718281828

	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 2 + 3i

	var stringVar string = "Hello, Go!"

	var boolVar bool = true

	var runeVar rune = 'A'

	var pointerVar *int = nil

	fmt.Println("int:", intVar)
	fmt.Println("int8:", int8Var)
	fmt.Println("int16:", int16Var)
	fmt.Println("int32:", int32Var)
	fmt.Println("int64:", int64Var)
	fmt.Println("uint:", uintVar)
	fmt.Println("uint8:", uint8Var)
	fmt.Println("uint16:", uint16Var)
	fmt.Println("uint32:", uint32Var)
	fmt.Println("uint64:", uint64Var)
	fmt.Println("float32:", float32Var)
	fmt.Println("float64:", float64Var)
	fmt.Println("complex64:", complex64Var)
	fmt.Println("complex128:", complex128Var)
	fmt.Println("string:", stringVar)
	fmt.Println("bool:", boolVar)
	fmt.Printf("rune: %c\n", runeVar)
	fmt.Println("pointer:", pointerVar)
}