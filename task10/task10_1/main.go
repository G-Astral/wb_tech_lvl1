// Дана последовательность температурных колебаний:
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"sort"
)

func getTens(n float64) int {
	return int(n/10.0) * 10
}

func main() {
	var keyArr []int

	valueArr := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	resMap := make(map[int][]float64)
	uniqueMap := make(map[int]struct{})

	// Вычленить ключи и засунуть в массив с ключами без повторений
	for _, v := range valueArr {
		uniqueMap[getTens(v)] = struct{}{}
	}

	for k := range uniqueMap {
		keyArr = append(keyArr, k)
	}

	sort.Ints(keyArr)

	// Добавить выходной мапе ключи
	for _, v := range keyArr {
		resMap[v] = []float64{}
	}

	// Сравнить значения в массиве с ключами и засунуть в мапу значения
	for _, i := range valueArr {
		key := getTens(i)
		resMap[key] = append(resMap[key], i)
	}

	fmt.Println(resMap)
}
