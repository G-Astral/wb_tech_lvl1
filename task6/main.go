// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"fmt"
	"sync"
)

func addText(n int) string {
	return fmt.Sprintf("ЗДЕСЬ РАБОТАЕТ ГОРУТИНА НОМЕР: %d", n)
}

func worker(jobs chan int, result chan string, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job := <-jobs:
			result <- addText(job)
		case <-done:
			return
		}
	}
}

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
		go worker(job, result, done, &wg)
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
