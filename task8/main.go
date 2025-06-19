// Дана переменная int64. Разработать программу
// которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func changBit(num int64, i, f int) int64 {
	switch f {
	case 0:
		num = num &^ (1 << i)
	case 1:
		num = num | (1 << i)
	}

	return num
}

func main() {
	var i int
	var flag int
	var number int64

	fmt.Print("Введите номер бита, который хоите изменить: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Ошибка при вводе данных")
		return
	}

	fmt.Print("Выбранный бит изменяется на [1/0] -> ")
	_, err1 := fmt.Scan(&flag)
	if err1 != nil || (flag != 0 && flag != 1) {
		fmt.Println("Ошибка при вводе данных")
		return
	}

	number = 8

	fmt.Println("До:    ", number)
	fmt.Println("После: ", changBit(number, i, flag))
}
