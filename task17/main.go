// Реализовать бинарный поиск
// встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{3, 5, 7, 9, 10}
	target := 10
	index := sort.SearchInts(arr1, target)

	if index < len(arr1) && arr1[index] == target {
		fmt.Println("Индекс найденного элемента:", index)
	} else {
		fmt.Println("Элемент не найден")
	}
}
