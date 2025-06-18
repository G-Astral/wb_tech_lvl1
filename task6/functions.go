package main

import (
	"context"
	"fmt"
	"sync"
)

func addText(n int) string {
	return fmt.Sprintf("ЗДЕСЬ РАБОТАЕТ ГОРУТИНА НОМЕР: %d", n)
}

func worker1(jobs chan int, result chan string, done chan struct{}, wg *sync.WaitGroup) {
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

func worker2(jobs chan int, result chan string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case job := <-jobs:
			result <- addText(job)
		case <-ctx.Done():
			return
		}
	}
}

func worker3(jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println(addText(job))
	}
}
