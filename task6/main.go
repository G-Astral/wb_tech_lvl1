// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Способы остановки горутин")
	fmt.Println("=========================")
	fmt.Println("\nЧерез done-канал :")

	job := make(chan int)
	result := make(chan string)
	done := make(chan struct{})
	workerCounter := 5
	var wg sync.WaitGroup

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker1(job, result, done, &wg)
	}

	go func() {
		for i := range result {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 3; i++ {
		job <- i + 1
	}

	close(done)
	wg.Wait()
	close(result)

	fmt.Println("Способ завершился.")
	main2()
}
