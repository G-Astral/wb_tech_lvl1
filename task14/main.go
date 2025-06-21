// Разработать программу, которая в рантайме
// способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}

package main

import "fmt"

func main()  {
	var any interface{}

	any = 123
	fmt.Printf("%v	is type of:	%T\n", any, any)
	
	any = "hello"
	fmt.Printf("%v	is type of:	%T\n", any, any)
	
	any = true
	fmt.Printf("%v	is type of:	%T\n", any, any)
	
	any = make(chan int, 1)
	ch := any.(chan int)
	ch <- 5
	value := <- ch
	fmt.Printf("%v	is type of:	%T", value, any)
}