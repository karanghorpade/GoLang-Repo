package main

import (
	"fmt"
)

type circle struct {
	radius int
}

type square struct {
	side int
}

type shape struct {
	circle
	square
}

func main() {

	c1 := circle{

		radius: 20,
	}

	s1 := square{
		side: 10,
	}

	sh1 := shape{

		circle: circle{
			radius: 10,
		},
		square: square{
			side: 5,
		},
	}

	a := c1.circlearea()
	fmt.Println("Area of circle:", a)
	//fmt.Println("Area of circle:", *a)

	b := s1.squarearea()
	fmt.Println("Area of square:", b)

	c, d := sh1.info()
	fmt.Println("The area of circle is:", c, "The area of square is:", d)

}

func (c circle) circlearea() float64 {
	pi := 3.14
	var area float64
	area = pi * float64(c.radius*c.radius)

	return area

}

func (s square) squarearea() int {

	var area int
	area = s.side * s.side
	return area
}

func (sh shape) info() (float64, int) {

	pi := 3.14
	var carea float64
	carea = pi * float64(sh.radius*sh.radius)

	var sarea int
	sarea = sh.side * sh.side
	return carea, sarea
}