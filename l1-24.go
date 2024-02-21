package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{4, 6}
	fmt.Println(Distance(p1, p2))
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(((p2.x - p1.x) * (p2.x - p1.x)) + ((p2.y - p1.y) * (p2.y - p1.y)))
}
