// Реализовать паттерн «адаптер» на любом примере

package main

import (
	"fmt"
)

// Target interface — твой ожидаемый интерфейс (написал гпт)
type Logger interface {
	Log(message string)
}

// Adaptee — чужой, несовместимый интерфейс (написал гпт)
type FileWriter struct{}

func (fw *FileWriter) WriteToFile(data []byte) {
	fmt.Println("Записываю в файл:", string(data))
}

// Adapter — обёртка, которая делает перевод между ними (написал я)
type FileWriterAdapter struct {
	fw FileWriter
}

func (fwa *FileWriterAdapter) Log(m string) {
	var d []byte = []byte(m)

	fwa.fw.WriteToFile(d)
}

func main() {
	var fwa FileWriterAdapter
	msg := "РАБОТАЕТ"

	fwa.Log(msg)
}
