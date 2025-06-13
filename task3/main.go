package main

import (
	"fmt"
	"sync"
)

func getSquare(n int) int {
	return n*n
}

// func getSum(sum *int, ch <-chan int) {
// 	*sum += <- ch
// }

func main()  {
	// arr := [5]int{2, 4, 6, 8, 10}
	// ch := make(chan int)
	// sum := 0
	
	// for _, v := range arr {
	// 	go func(n int) {
	// 		ch <- getSquare(n)
	// 	}(v)
	// }

	// for i := 0; i < len(arr); i++ {
	// 	sum += <- ch
	// }
		
	// // go getSum(&sum, ch)

	// fmt.Println(sum)

	// main2()
	main3()
}

// func getSquare2(n int, ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ch <- n * n
// }


// func main2() {
// 	arr := [5]int{2, 4, 6, 8, 10}
// 	ch := make(chan int)
// 	sum := 0
// 	var wg sync.WaitGroup

// 	for _, v := range arr {
// 		wg.Add(1)
// 		go func() {
// 			getSquare2(v, ch, &wg)
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for c := range ch {
// 		sum += c
// 	}

// 	fmt.Println(sum)

// }


func worker(jobs, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result <- getSquare(job)
	}
}


func main3() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	job := make(chan int, 5)
	result := make(chan int, 5)
	var wg sync.WaitGroup

	countWorker := 5

	// start worker
	for i := 0; i < countWorker; i++ {
		wg.Add(1) 
		go worker(job, result, &wg)
	}


	// add job
	go func() {
		for _, i := range arr {
			job <- i
		}
		close(job)
	}()

	// wait workers
	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		sum += res
	}

	fmt.Println(sum)
}