package main

import ("fmt")

func getSquare(n int) int {
	return n*n
}

// func getSum(sum *int, ch <-chan int) {
// 	*sum += <- ch
// }

func main()  {
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	sum := 0
	
	for _, v := range arr {
		go func(n int) {
			ch <- getSquare(n)
			}(v)
			sum += <- ch
		}
		
	// go getSum(&sum, ch)

	fmt.Println(sum)
}