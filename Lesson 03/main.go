package main

import "fmt"

type rectangle struct {
	Width  int
	Length int
}

func (r rectangle) Area() int {
	return r.Length * r.Width
}

func (r rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

func (r *rectangle) Scale2X() {
	r.Length = r.Length * 2
	r.Width = r.Width * 2
}

func main() {
	rect := rectangle{Length: 4, Width: 5}
	fmt.Printf("Retangle: %+v\n", rect)
	fmt.Printf("Area: %d\n", rect.Area())
	fmt.Printf("Perimeter: %d\n", rect.Perimeter())

	rect.Scale2X()

	fmt.Printf("Retangle: %+v\n", rect)
}
