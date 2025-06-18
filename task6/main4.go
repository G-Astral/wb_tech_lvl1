// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
)

func main4() {
	fmt.Println("\nЧерез select и закрытие канала:")

	fmt.Println("Способ завершился.")
	main5()
}
