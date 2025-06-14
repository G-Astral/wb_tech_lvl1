// Разработать программу, которая будет последовательно отправлять
// значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func printer(s string) {
	fmt.Printf("Было прочтено: %s\n", s)
}

func worker(jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		printer(job)
	}
}

func main() {
	var n int
	fmt.Print("Введите время действия программы в секундах: ")
	_, err1 := fmt.Scan(&n)
	if err1 != nil {
		fmt.Println("Ошибка в вводе данных")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	job := make(chan string)
	var wg sync.WaitGroup
	workerCounter := 10

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker(job, &wg)
	}

	go func() {
		wg.Wait()
	}()

	go func() {
		for i := range job {
			fmt.Println(i)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			close(job)
			return
		default:
			if scanner.Scan() {
				job <- scanner.Text()
			}
		}
	}
}
