// Реализовать все возможные способы
// остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main2() {
	fmt.Println("\nЧерез contex:")

	job := make(chan int)
	result := make(chan string)
	var wg sync.WaitGroup
	workerCounter := 5

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker2(job, result, &wg, ctx)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for i := 0; i < 3; i++ {
					job <- i + 1
				}
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()

	go func() {
		for i := range result {
			fmt.Println(i)
		}
	}()

	wg.Wait()
	close(job)
	close(result)

	fmt.Println("Способ завершился.")
	main3()
}
