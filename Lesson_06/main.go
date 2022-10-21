package main

import (
	"fmt"
	"log"
)

type MinInputError struct {
}

func (e *MinInputError) Error() string {
	return "invalid input, number must be greater than 0"
}

type MaxInputError struct {
}

func (e *MaxInputError) Error() string {
	return "invalid input, number must be less than 100"
}

func Factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}

func Factorial2(n int) (int, error) {
	if n < 0 {
		return 0, &MinInputError{}
	}
	if n > 100 {
		return 0, &MaxInputError{}
	}
	ret := n
	for i := n - 1; i > 0; i-- {
		ret = ret * i
	}
	return ret, nil
}

func main() {
	fmt.Println(Factorial(10))
	fmt.Println(Factorial2(10))
	fmt.Println(Factorial2(-10))

	f, err := Factorial2(-13)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(f)
	}

	f, err = Factorial2(100000000000)

	if _, ok := err.(*MaxInputError); ok {
		fmt.Println("Do something special")
	} else {
		fmt.Println(f)
	}

}
