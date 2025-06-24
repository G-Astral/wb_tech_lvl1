// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func mySleep(d float64, t string) {
	switch t {
	case "Nanosecond":
		d *= 1e-9
	case "Microsecond":
		d *= 1e-6
	case "Millisecond":
		d *= 1e-3
	case "Second":
		d *= 1
	case "Minute":
		d *= 60
	case "Hour":
		d *= 3600
	default:
		fmt.Println("Неверный второй аргумент!")
		return
	}

	fmt.Printf("Ложимся спать на %f секунд\n", d)
	<-time.After(time.Duration(d) * time.Second)
	fmt.Println("Проснулись, улыбнулись")
}

func main() {
	n := 5.0

	mySleep(n, "Second")
}
