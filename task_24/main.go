/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}
func (p Point) Distance(q Point) float64 {
	return math.Sqrt(((q.x-p.x)*(q.x-p.x) + (q.y-p.y)*(q.y-p.y)))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	fmt.Println(p1.Distance(p2))
}
