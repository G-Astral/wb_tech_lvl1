// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

type Pair struct {
	Key   string
	Value string
}

func worker(jobs chan Pair, wg *sync.WaitGroup, res map[string]string, mu *sync.Mutex) {
	defer wg.Done()

	for job := range jobs {
		mu.Lock()
		res[job.Key] = job.Value
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	keyArr := [10]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth"}
	valueArr := [10]string{"Poshel", "uchit", "go", "teper", "v", "production", "i", "rabotay", "kak", "boh"}

	job := make(chan Pair)
	result := make(map[string]string)
	workerCounter := 5

	for i := 0; i < workerCounter; i++ {
		wg.Add(1)
		go worker(job, &wg, result, &mu)
	}

	for i := 0; i < 10; i++ {
		job <- Pair{Key: keyArr[i], Value: valueArr[i]}
	}

	close(job)

	wg.Wait()

	for k, v := range result {
		fmt.Printf("%s: %s\n", k, v)
	}
}
