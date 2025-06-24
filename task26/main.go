// Разработать программу, которая проверяет,
// что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	tempMap := make(map[rune]struct{})

	for _, v := range s {

		if _, exists := tempMap[v]; exists {
			return false
		} else {
			tempMap[v] = struct{}{}
		}
	}

	return true
}

func main() {
	s := "esd3ghf"
	s = strings.ToLower(s)

	fmt.Print(s, " - ", checkUnique(s))
}
