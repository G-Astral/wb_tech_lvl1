// Разработать программу, которая
// переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку:  ")

	if scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}

	res := []rune(input)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	fmt.Println("Результат:	", string(res))
}
