// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Способы остановки горутин")
	fmt.Println("=========================")
	fmt.Println("\nЧерез закрытие канала:")

	ch1 := make(chan string, 5)

	go func() {
		ch1 <- "ТУТ РАБОТАЕТ ГОРУТИНА"
		ch1 <- "ТУТ ГОРУТИНА РАБОТАТЬ НЕ ДОЛЖНА"
		fmt.Println(<-ch1)
		close(ch1)
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Способ завершился.")
	main2()
}
