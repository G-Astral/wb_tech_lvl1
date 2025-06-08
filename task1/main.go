package main

import (
	// "fmt"
	// "strings"
)

type Human struct {
	name string
	age int
}

type Action struct {
	Human
}

func (a *Action) changeAge(n int) {
	a.Human.age = n
}

func (a *Action) changeName(s string) {
	a.Human.name = s
}

// func main() {
// 	Human.name = "Gar"
// 	Human.age = 23

// }