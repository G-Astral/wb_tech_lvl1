package main

import (
	"fmt"
	"sync"
)

// Инструментарий работника
func getSquare(n int) int {
	return n*n
}

// Сам работник
func worker(jobs, result chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Закрыли задачу

	// Работник выполняет задачу из канала задач и добавляет её в канал результатов
	for job := range jobs {
		result <- getSquare(job)
	}
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup // Счётчик задач
	job := make(chan int) // Канал задач размером 5
	result := make(chan int, 5) // Канал результатов размером 5
	workerCounter := 5 // Максимальное количество одновременно работающих работников

	// Начать работу воркера
	for i := 0; i < workerCounter; i++ {
		wg.Add(1) // Открыли задачу
		go worker(job, result, &wg) // Начали работу работника
	}

	// Добавить воркеру задачу
	go func() {
		for _, i := range arr {
			job <- i
		}
		close(job) // Задача выполнена - надо её закрыть		
	}()

	// Забрать у воркера результат задачи и закончить его работу
	go func() {
		wg.Wait()
		close(result)
	}()

	// Суммировать результат работы воркеров за весь день (именно в этом задании
	// for i := 0; i < len(arr); i++ {
	for i := range result {
		 sum += i
	}

	// Вывод результата "рабочего дня"
	fmt.Println(sum)
}