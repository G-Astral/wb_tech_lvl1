// К каким негативным последствиям может привести
// данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }


// Проблема в том, что переменная объявлена вне мэйна
// Происходит утечка памяти 
// Нужно сделать полную копию среза строки
// Ниже правильная реализация
package main

// import "fmt"

// func someFunc(s string) string{
// 	v := createHugeString(1 << 10)
// 	s = string([]byte(v[:100]))
// 	return s
// }

// func main() {
// 	var justString string
// 	someFunc(justString)

// 	fmt.Println(justString)
// }