package main

import (
	"fmt"
	"math"
)

/*
type Circle struct {
	x float64
	y float64
	z float64
}
*/

type Circle struct {
	x, y, z float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	//var c Circle
	// c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	//c := Circle{0, 0, 5}
	c := &Circle{0, 0, 5}

}
