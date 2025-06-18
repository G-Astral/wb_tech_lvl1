// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
	"sync"
)

func main3() {
	fmt.Println("\nЧерез закрытие канала:")

	job := make(chan int)
	var wg sync.WaitGroup
	workerCounter := 5

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker3(job, &wg)
	}

	for i := 0; i < 3; i++ {
		job <- i + 1
	}

	close(job)

	wg.Wait()

	fmt.Println("Способ завершился.")
}
