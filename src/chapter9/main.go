package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Shape ... could be any shape like Rectangle or circle
type Shape interface {
	area() float64
	perimeter() float64
}

// Rectangle ... should implement area and perimeter
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (rect *Rectangle) area() float64 {
	l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	w := distance(rect.x1, rect.y1, rect.x2, rect.y1)
	return l * w
}

func (rect *Rectangle) perimeter() float64 {
	l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	w := distance(rect.x1, rect.y1, rect.x2, rect.y1)
	return 2*l + 2*w
}

// Circle ... center cordinates and radius
type Circle struct {
	// x, y and radius
	x, y, r float64
}

func (circle *Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

func (circle *Circle) perimeter() float64 {
	return math.Pi * circle.r * 2
}

// MultiShape ... a list of shapes
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var p float64
	for _, s := range m.shapes {
		p += s.perimeter()
	}
	return p
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{0,0,5}
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r, c.area())
	fmt.Println(c.x, c.y, c.r, c.perimeter())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	m := MultiShape{shapes: []Shape{&c, &r}}
	fmt.Println(m.area())
	fmt.Println(m.perimeter())
}
