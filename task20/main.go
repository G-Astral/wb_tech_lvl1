// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	var arr []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ввод:  ")

	if scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Воззникла ошибка при чтении:", err)
	}

	arr = strings.Split(line, " ")

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("Вывод:", strings.Join(arr, " "))
}
