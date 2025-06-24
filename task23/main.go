// Удалить i-ый элемент из слайса

package main

import "fmt"

func remove(a []int, n int) []int {
	return append(a[:n], a[n+1:]...)
}

func main() {
	var n int
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr)
	fmt.Print("Введите, какой элемент слайса вы хотите удалить (начиная с 0): ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Ошибка при чтении данных")
		return
	}

	if n < 0 || n >= len(arr) {
		fmt.Println("Недопустимое значение")
		return
	}

	fmt.Println(remove(arr, n))
}
