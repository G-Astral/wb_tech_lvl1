// Разработать программу, которая
// перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"math"
)

func add(a, b int64) int64 {
	return a + b
}

func substract(a, b int64) int64 {
	return a - b
}

func multiply(a, b int64) int64 {
	return a * b
}

func divide(a, b int64) int64 {
	return a / b
}

func main() {
	var a, b int64

	a, b = 1523456789876594864, 4765341

	if a < int64(math.Pow(2, 20)) || b < int64(math.Pow(2, 20)) {
		fmt.Println("Слишком маленькие числа")
		return
	}

	fmt.Println("Сложение:	", add(a, b))
	fmt.Println("Вычитание:	", substract(a, b))
	fmt.Println("Умножение:	", multiply(a, b))
	fmt.Println("Деление:	", divide(a, b))
}
