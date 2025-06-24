// Разработать программу нахождения расстояния
// между двумя точками, которые представлены в виде
// структуры Point с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) newPoint(x1, y1 float64) {
	p.x = x1
	p.y = y1
}

func getDistance(a, b Point) float64 {
	return math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
}

func main() {
	var a, b Point

	a.newPoint(5, 10)
	b.newPoint(-3, 9)

	fmt.Println(getDistance(a, b))
}
