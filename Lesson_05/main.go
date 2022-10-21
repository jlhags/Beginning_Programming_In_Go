package main

import (
	"fmt"
	"math"
)

type Polygon interface {
	Shape() string
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Length float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Length
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Length
}
func (r Rectangle) Shape() string {
	return "rectangle"
}
func (r Rectangle) String() string {
	return fmt.Sprintf("Shape: %s, Dimensions: %fx%f, Perimeter: %f, Area: %f", r.Shape(), r.Length, r.Height, r.Perimeter(), r.Area())
}

type ColoredRectangle struct {
	Rectangle
	Color string
}

func (r ColoredRectangle) String() string {
	return fmt.Sprintf("Shape: %s, Dimensions: %fx%f, Perimeter: %f, Area: %f, Color: %s", r.Shape(), r.Length, r.Height, r.Perimeter(), r.Area(), r.Color)
}

type Triangle struct {
	Sides [3]float64
}

func (t Triangle) Perimeter() float64 {
	return t.Sides[0] + t.Sides[1] + t.Sides[2]
}

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.Sides[0]) * (p - t.Sides[1]) * (p - t.Sides[2]))
}

func (t Triangle) Shape() string {
	return "triangle"
}

func (t Triangle) String() string {
	return fmt.Sprintf("Shape: %s, Dimensions: %fx%fx%f, Perimeter: %f, Area: %f", t.Shape(), t.Sides[0], t.Sides[1], t.Sides[2], t.Perimeter(), t.Area())
}

func printSlice(s []interface{}) {
	for i, v := range s {
		fmt.Printf("%d - %v\n", i, v)
	}
}

func main() {
	rect := Rectangle{Length: 5, Height: 4}
	tri := Triangle{Sides: [3]float64{4, 5, 6}}

	fmt.Println(rect)

	fmt.Println(tri)

	var polys []Polygon

	polys = append(polys, rect)
	polys = append(polys, tri)

	var cr ColoredRectangle
	cr.Height = 7
	cr.Length = 8
	cr.Color = "red"

	polys = append(polys, cr)

	fmt.Println(polys)

	for _, p := range polys {
		fmt.Println(p)
	}

	fmt.Printf("Shape 1 type: %s\n", polys[0].Shape())
	random := []interface{}{1, "blue", rect, tri}

	printSlice(random)

	if _, ok := random[1].(Polygon); !ok {
		fmt.Println("Not a polygon")
	}

	fmt.Println(random[2].(Polygon).Shape())

}
