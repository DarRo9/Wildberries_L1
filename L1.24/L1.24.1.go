//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 
}

func NewPoint(x float64, y float64) *Point { 
	return &Point{x: x, y: y}
}

func LineLength(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	point1 := NewPoint(7.2, 5.4)
	point2 := NewPoint(8.9, 3.7)
	fmt.Println(LineLength(point1, point2))
}
