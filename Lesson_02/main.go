package main

import "fmt"

func doubleFloat(f float64) float64 {
	return f * 2
}

func doubleInt(n *int) {
	(*n) = (*n) * 2
}

func multiply(a int, b int) int {
	return a * b
}

func multiReturn(i int) (int, int) {
	return i - 1, i + 1
}

func main() {
	//number types
	var i int
	var f float64 = 2

	//text types
	var c rune
	var s string

	// boolean
	var b bool

	i = 12
	f = 1.41

	c = 'r'
	s = "Here is some text"
	fmt.Print("i: ")
	fmt.Println(i)

	fmt.Print("f: ")
	fmt.Println(f)

	fmt.Print("c: ")
	fmt.Println(c)
	fmt.Printf("c: %c\n", c)

	fmt.Print("s: ")
	fmt.Println(s)

	fmt.Print("b: ")
	fmt.Println(b)

	fmt.Print("f * 2: ")
	fmt.Println(doubleFloat(f))

	fmt.Print("f: ")
	fmt.Println(f)

	doubleInt(&i)
	fmt.Print("i: ")
	fmt.Println(i)

	fmt.Println(multiply(2, 3))

	i1, i2 := multiReturn(i)
	fmt.Printf("%d:%d", i1, i2)

}
