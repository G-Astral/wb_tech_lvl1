package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// func printText(s string) string {
// 	return fmt.Sprintf("Я ВЫВЕЛ ТЕКСТ: %s\n", s)
// }

func printText(s string) string {
	return fmt.Sprintf("Я ВЫВЕЛ ТЕКСТ: %s\n\nВвод текста: ", s)
}

func worker(jobs, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result <- printText(job)
	}
}

func main() {
	// +принимаешь кол-во воркеров
	var workerCounter int
	fmt.Print("Введите количество воркеров N = ")
	_, err1 := fmt.Scan(&workerCounter)

	if err1 != nil {
		fmt.Println("Ошибка при записи данных")
	}

	// +иниц каналы: job, res
	var wg sync.WaitGroup
	job := make(chan string)
	result := make(chan string)

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker(job, result, &wg)
	}

	go func() {
		wg.Wait()
		close(job)
		close(result)
	}()

	go func() {
		for i := range result {
			fmt.Println(i)
		}
	}()

	// запуск job
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		job <- scanner.Text()
	}

	// запуск чтение в отд горутине res -> вывод переменной
	// запуск сбор данных в горутине -> job канал -> res

}

// func main2() {
// 	// +принимаешь кол-во воркеров
// 	var workerCounter int
// 	fmt.Print("Введите количество воркеров N = ")
// 	_, err1 := fmt.Scan(&workerCounter)

// 	if err1 != nil {
// 		fmt.Println("Ошибка при записи данных")
// 	}

// 	// +иниц каналы: job, res
// 	var wg sync.WaitGroup
// 	job := make(chan string)
// 	result := make(chan string)

// 	for i := 0; i < workerCounter; i++ {
// 		wg.Add(1)
// 		go worker(job, result, &wg)
// 	}

// 	// запуск job
// 	go func() {
// 		defer close(job)
// 		scanner := bufio.NewScanner(os.Stdin)
// 		for {
// 			fmt.Print("Ввод значения: ")
// 			scanner.Scan()
// 			job <- scanner.Text()

// 		}

// 	}()

// 	// запуск чтение в отд горутине res -> вывод переменной
// 	// запуск сбор данных в горутине -> job канал -> res

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	for i := range result {
// 		fmt.Println(i)
// 	}
// }

// func main3(){
// 	var workerCounter int
// 	fmt.Print("Введите количество воркеров N = ")
// 	_, err1 := fmt.Scan(&workerCounter)

// 	if err1 != nil {
// 		fmt.Println("Ошибка при записи данных")
// 	}

// 	var stringArr []string

// 	enterValue(&stringArr)

// 	var wg sync.WaitGroup
// 	job := make(chan string)
// 	result := make(chan string)

// 	for i := 0; i < workerCounter; i++ {
// 		wg.Add(1)
// 		go worker(job, result, &wg)
// 	}

// 	go func() {
// 		for _, i := range stringArr {
// 			job <- i
// 		}
// 		close(job)
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	for i := range result {
// 		fmt.Println(i)
// 	}
// }

// func enterValue(sArr *[]string)  {
// 	var s string
// 	for {
// 	// for i := 0; i < 5; i++ {
// 		fmt.Print("Ввод значения: ")
// 		_, err2 := fmt.Scan(&s)

// 		if err2 != nil {
// 			fmt.Println("Ошибка при записи данных")
// 		}

// 		*sArr = append(*sArr, s)
// 	}
// }
