// Поменять местами два числа
// без создания временной переменной

package main

import "fmt"

func main()  {
	num1 := 9
	num2 := 6

	fmt.Println("Было:  ", num1, num2)
	
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	
	fmt.Println("Стало: ", num1, num2)
}