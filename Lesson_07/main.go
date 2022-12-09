package main

import (
	"fmt"

	"github.com/jlhags/tutorials/lesson_07/polygons"
)

func main() {
	rect := polygons.Rectangle{Length: 5, Height: 4}
	tri := polygons.Triangle{Sides: [3]float64{4, 5, 6}}

	fmt.Println(rect)

	fmt.Println(tri)
	rect.Area()
	//rect.internalUseOnly()
}
