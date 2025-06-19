// Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func changeNum(n int) int {
	n *= 2
	return n
}

func worker(jobs, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result <- changeNum(job)
	}
}

func main() {
	var wg sync.WaitGroup

	job := make(chan int)
	result := make(chan int)
	workerCounter := 5
	numArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker(job, result, &wg)
	}

	go func() {
		for i := range result {
			fmt.Println(i)
		}
	}()

	for _, i := range numArr {
		job <- i
	}

	close(job)
	wg.Wait()
	close(result)
}
