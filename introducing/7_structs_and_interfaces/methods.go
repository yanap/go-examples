package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, z float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	//Person Person
	Person
	Model string
}


func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.x * c.y
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	c := new(Circle)
	fmt.Println(c.area)
	a := new(Android)
	a.Talk()
}

type Shape interface {
	area() float64
}

func totalArea(circles ...Circle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	return total
}
