package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{4, 5, 6}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s1 Length: %d\n", len(s1))
	fmt.Printf("s2 Length: %d\n", len(s2))

	s1 = append(s1, 1)
	fmt.Printf("s1 Length: %d\n", len(s1))
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Println(s1)

	fmt.Println(s1[1])
	fmt.Println(s1[0])
	fmt.Printf("Max s2: %d\n", max(s2))

	p := polygon{Sides: []float64{1, 1, 1.4142}}

	fmt.Printf("Perimeter: %f\n", p.Perimeter())
	fmt.Println(p.Type())

	p2 := polygon{Sides: []float64{1, 1, 1, 1, 1, 1, 1}}
	fmt.Println(p2.Type())
}

func max(s []int) int {
	var ret int
	// for i := 0; i < len(s); i++ {
	// 	if s[i] > ret {
	// 		ret = s[i]
	// 	}
	// }
	for _, v := range s {
		if v > ret {
			ret = v
		}
	}
	return ret
}

type polygon struct {
	Sides []float64
}

func (p polygon) Perimeter() float64 {
	var ret float64

	for _, v := range p.Sides {
		ret += v
	}
	return ret
}

func (p polygon) IsRegular() bool {
	ret := true

	if len(p.Sides) < 3 {
		ret = false
	} else {
		for i := 1; i < len(p.Sides); i++ {
			if p.Sides[i-1] != p.Sides[i] {
				ret = false
				break
			}
		}
	}

	return ret
}

func (p polygon) Type() string {
	var ret string
	switch len(p.Sides) {
	case 3:
		ret = "triangle"
	case 4:
		ret = "quadrangle"
	case 5:
		ret = "pentagon"
	default:
		ret = "unknown"
	}
	return ret
}
