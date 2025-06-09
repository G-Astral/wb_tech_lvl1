package main

import ("fmt")

func getSquare(n int) int {
	return n*n
}

func main()  {
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for _, v := range arr {
		go func(n int) {
			ch <- getSquare(n)
		}(v)
	}

	for i := 0; i < len(arr); i++ {
		
		fmt.Println(<-ch)
	}
}