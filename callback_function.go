package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(a...)

	e := even(sum, a...)
	fmt.Println(s)
	fmt.Println("Sum of even nos:",e)
}

func sum(xi ...int) int {

	total := 0
	for _, v := range xi {

		total += v
	}
	return total
}

func even(f func(xi ...int) int, b ...int) int {

	var c []int
	for _, v := range b {

		if v%2 == 0 {

			c = append(c, v)
		}
	}

	return f(c...)
}

