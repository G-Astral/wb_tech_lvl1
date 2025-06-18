// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
)

func main5() {
	fmt.Println("\nЧерез флаг + sync/atomic:")

	fmt.Println("Способ завершился.")
	main6()
}
