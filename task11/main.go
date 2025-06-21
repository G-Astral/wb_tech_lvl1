// Реализовать пересечение двух неупорядоченных множеств

package main

import (
	"fmt"
	"sort"
)

func main() {
	var resArr []int
	arrA := []int{
		12, 57, 3, 89, 34, 76, 45, 18, 99, 60,
		23, 8, 77, 42, 6, 91, 29, 55, 80, 14,
		41, 73, 27, 92, 63, 5, 36, 81, 50, 19,
		2, 70, 84, 10, 65, 47, 31, 17, 90, 40,
		13, 38, 22, 74, 59, 67, 11, 21, 79, 33,
	}

	arrB := []int{
		33, 76, 55, 14, 88, 90, 3, 21, 42, 68,
		99, 2, 6, 12, 35, 70, 50, 18, 77, 9,
		41, 63, 25, 92, 65, 17, 5, 79, 24, 57,
		84, 8, 36, 40, 22, 47, 73, 10, 19, 31,
		38, 91, 60, 59, 13, 29, 27, 11, 80, 74,
	}

	tempMap := make(map[int]struct{})

	for _, v := range arrA {
		tempMap[v] = struct{}{}
	}

	for _, v := range arrB {
		if _, exists := tempMap[v]; exists {
			resArr = append(resArr, v)
		}
	}

	sort.Ints(resArr)

	fmt.Println(resArr)
}