package main

import "fmt"

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

func main() {
	a := Action{}
	a.name = "Gar"
	a.age = 23
	fmt.Println(a.name, a.age)
	
	a.changeName("Sergey")
	a.changeAge(24)
	fmt.Println(a.name, a.age)
}