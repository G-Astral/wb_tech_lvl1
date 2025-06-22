// Реализовать структуру-счетчик, которая будет
// инкрементироваться в конкурентной среде.
// По завершению программа должна выводить
// итоговое значение счетчика

package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func numChanger(n int) int {
	return n * 2
}

func worker(jobs, result chan int, wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()

	for job := range jobs {
		result <- numChanger(job)
		c.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	var count Counter

	job := make(chan int)
	result := make(chan int)
	workerCounter := 1
	arr1 := []int{
		345, 728, 99, 532, 413, 872, 156, 274, 617, 901,
		231, 84, 767, 620, 413, 257, 945, 372, 81, 523,
	}
	sort.Ints(arr1)
	fmt.Println(arr1)

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker(job, result, &wg, &count)
	}

	go func() {
		for i := range result {
			fmt.Print(i, " ")
		}
	}()

	for _, i := range arr1 {
		job <- i
	}

	close(job)
	wg.Wait()
	close(result)

	fmt.Println("\nРезультат счетчика:", count.value)
}
